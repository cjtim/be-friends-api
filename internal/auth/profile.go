package auth

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cjtim/be-friends-api/configs"
)

type LineProfile struct {
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

func getProfile(token string) (LineProfile, error) {
	// Get Data
	resp, err := http.PostForm("https://api.line.me/oauth2/v2.1/verify", url.Values{
		"id_token":  {token},
		"client_id": {configs.Config.LineClientID},
	})
	if err != nil {
		return LineProfile{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return LineProfile{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return LineProfile{}, errors.New(string(body))
	}

	// Data
	userInfo := LineToken{}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return LineProfile{}, err
	}

	profile := LineProfile{}
	err = json.Unmarshal(body, &profile)
	return profile, err
}
