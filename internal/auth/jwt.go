package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/utils"
	"github.com/cjtim/be-friends-api/repository"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var (
	TOKEN_EXPIRE = time.Hour * 72
)

type LineToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func GetNewToken(u *repository.User) (*jwt.Token, string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":          u.ID,
		"name":        u.Name,
		"email":       u.Email,
		"line_uid":    u.LineUid,
		"picture_url": u.PictureURL,
		"exp":         time.Now().Add(TOKEN_EXPIRE).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(configs.Config.JWTSecret))
	if err != nil {
		return token, "", err
	}
	err = repository.RedisJwt.AddJwt(t, TOKEN_EXPIRE)
	return token, t, err
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
		zap.L().Error("error redis - cannot save callback", zap.Error(err))
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

	userInfo := LineToken{}
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
		return "", err
	}

	profile, err := getProfile(token)
	if err != nil {
		return "", err
	}

	// 2. Update database
	userDB, err := profile.createLineUser()
	if err != nil {
		return "", err
	}

	// 3. Create JWT
	userInfo := repository.User{
		ID:         userDB.ID,
		Name:       profile.Name,
		Email:      userDB.Email,
		LineUid:    userDB.LineUid,
		PictureURL: userDB.PictureURL,
	}
	_, jwtToken, err := GetNewToken(&userInfo)
	return jwtToken, err
}
