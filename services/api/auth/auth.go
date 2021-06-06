package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var authCheck string

type AuthResponse struct {
	Token      string  `json:"token"`
	Message    string  `json:"message"`
	User       UserDTO `json:"user"`
	StatusCode int     `json:"status_code"`
}

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func init() {
	authCheck = fmt.Sprintf("http://%s:%s/check",
		os.Getenv("AUTH_HOST"),
		os.Getenv("AUTH_PORT"))
}

func CheckAuth(token string) (AuthResponse, error) {
	authResp := AuthResponse{}

	req, err := http.NewRequest("GET", authCheck, nil)

	if err != nil {
		authResp.StatusCode = req.Response.StatusCode
		return authResp, err
	}

	req.Header.Add("Authorization", token)

	client := http.DefaultClient

	response, err := client.Do(req)

	if err != nil {
		return authResp, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&authResp)
	if err != nil {
		return authResp, err
	}

	authResp.StatusCode = response.StatusCode

	if response.StatusCode == http.StatusForbidden {
		return authResp, fmt.Errorf("%s", authResp.Message)
	}

	authResp.StatusCode = response.StatusCode

	return authResp, nil
}
