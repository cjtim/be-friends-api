package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/utils"
	"github.com/cjtim/be-friends-api/repository"
	"go.uber.org/zap"
)

type lineToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type lineProfile struct {
	Iss     string   `json:"iss"`
	LineUid string   `json:"sub"` // lineUid
	Aud     string   `json:"aud"`
	Exp     int      `json:"exp"`
	Iat     int      `json:"iat"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
	Email   string   `json:"email"`
}

func getProfile(token string) (lineProfile, error) {
	// Get Data
	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/verify", url.Values{
		"id_token":  {token},
		"client_id": {configs.Config.LineClientID},
	})
	if err != nil {
		return lineProfile{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return lineProfile{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return lineProfile{}, errors.New(string(body))
	}

	// Data
	userInfo := lineToken{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return lineProfile{}, err
	}

	profile := lineProfile{}
	err = json.Unmarshal(body, &profile)
	return profile, err
}

func GetLoginURL(clientURL string) string {
	url := http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "access.line.me",
			Path:   "/oauth2/v2.1/authorize",
		},
	}

	state := utils.RandomSeq(10)
	err := repository.RedisCallback.AddCallback(state, clientURL)
	if err != nil {
		zap.L().Error("error redis - cannot save callback", zap.String("clientURL", clientURL), zap.Error(err))
		return ""
	}

	q := url.URL.Query()
	q.Add("state", state)
	q.Add("scope", "profile openid")
	q.Add("response_type", "code")
	q.Add("redirect_uri", configs.Config.LineLoginCallback)
	q.Add("client_id", configs.Config.LineClientID)
	url.URL.RawQuery = q.Encode()
	return url.URL.String()
}

func getJWT(code string) (string, error) {
	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/token", url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {configs.Config.LineLoginCallback},
		"client_id":     {configs.Config.LineClientID},
		"client_secret": {configs.Config.LineSecretID},
	})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	userInfo := lineToken{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return "", err
	}
	return userInfo.IDToken, err
}

func Callback(code string) (string, error) {
	// 1. Get profile from LINE
	token, err := getJWT(code)
	if err != nil {
		zap.L().Error("error line get jwt", zap.String("code", code), zap.Error(err))
		return "", err
	}

	profile, err := getProfile(token)
	if err != nil {
		zap.L().Error("error line get profile", zap.String("token", token), zap.Error(err))
		return "", err
	}

	// 2. Update database
	userDB, err := profile.createLineUser()
	if err != nil {
		zap.L().Error("error create user line", zap.Any("profile", profile), zap.Error(err))
		return "", err
	}

	// 3. Create JWT
	_, jwtToken, err := GetNewToken(userDB.ID)
	return jwtToken, err
}
