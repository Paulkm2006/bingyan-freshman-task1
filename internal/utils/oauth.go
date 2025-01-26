package utils

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAccessToken(code string) (string, error) {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", config.Config.Oauth.ClientID, config.Config.Oauth.ClientSecret, code)

	resp, err := http.Post(url, "application/json", nil)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return "", err
	}

	return result["access_token"].(string), nil

}

func GetOrCreateUser(accessToken string) (*dto.User, error) {
	url_user := "https://api.github.com/user"
	url_email := "https://api.github.com/user/emails"

	client := http.Client{}

	req, err := http.NewRequest("GET", url_user, nil)

	req.Header.Set("Authorization", "bearer "+accessToken)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var user dto.User
	user.Username = result["login"].(string)
	user.Nickname = result["login"].(string)

	req, err = http.NewRequest("GET", url_email, nil)

	req.Header.Set("Authorization", "bearer "+accessToken)
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		if v.(map[string]interface{})["primary"].(bool) {
			user.Email = v.(map[string]interface{})["email"].(string)
			break
		}
	}

	user.Oauth = true

	return &user, nil
}
