package line

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
	"github.com/cjtim/be-friends-api/internal/utils"
)

type LineToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func GetLoginURL() string {
	url := http.Request{
		URL: &url.URL{
			Scheme: "https",
			Host:   "access.line.me",
			Path:   "/oauth2/v2.1/authorize",
		},
	}
	q := url.URL.Query()
	q.Add("state", utils.RandomSeq(10))
	q.Add("scope", "profile openid")
	q.Add("response_type", "code")
	q.Add("redirect_uri", configs.Config.LineLoginCallback)
	q.Add("client_id", configs.Config.LineClientID)
	url.URL.RawQuery = q.Encode()
	return url.URL.String()
}

func GetJWT(code, state string) (string, error) {
	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/token", url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"state":         {state},
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
