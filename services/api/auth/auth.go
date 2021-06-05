package auth

import (
	"fmt"
	"net/http"
	"os"
)

var authCheck string

func init() {
	authCheck = fmt.Sprintf("http://%s:%s/%s/check",
		os.Getenv("AUTH_HOST"),
		os.Getenv("AUTH_PORT"),
		os.Getenv("AUTH_PREFIX"))
}

func CheckAuth(token string) error {

	response, err := http.Get(authCheck)

	if err != nil {
		return err
	}

	if response.Status == fmt.Sprint(http.StatusForbidden) {
		return fmt.Errorf("Invalid token")
	}

	return nil
}
