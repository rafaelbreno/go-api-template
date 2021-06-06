package auth

import (
	"fmt"
	"net/http"
	"os"
)

var authCheck string

type AuthResponse struct {
	Token      string  `json:"token"`
	User       UserDTO `json:"user"`
	StatusCode int
}

type UserDTO struct {
	ID       uint   `json:"user_id"`
	Username string `json:"username"`
}

func init() {
	authCheck = fmt.Sprintf("http://%s:%s/%s/check",
		os.Getenv("AUTH_HOST"),
		os.Getenv("AUTH_PORT"),
		os.Getenv("AUTH_PREFIX"))
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

	authResp.StatusCode = response.StatusCode

	if err != nil {
		return authResp, err
	}

	if response.Status == fmt.Sprint(http.StatusForbidden) {
		return authResp, fmt.Errorf("Invalid token")
	}

	authResp.StatusCode = response.StatusCode

	return authResp, nil
}
