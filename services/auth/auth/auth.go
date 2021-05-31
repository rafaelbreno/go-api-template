package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rafaelbreno/go-api-template/services/auth/entity"
)

type Wrapper struct {
	Secret          string
	Issuer          string
	ExpirationHours int
}

type Claim struct {
	Username string
	jwt.StandardClaims
}

type UserAuth struct {
	Token string `json:"token"`
}

func GetUserToken(u entity.User) (UserAuth, error) {
	var userAuth UserAuth

	jwtWrapper := Wrapper{
		Secret:          "super-secret",
		Issuer:          "AuthService",
		ExpirationHours: 48,
	}

	signedToken, err := jwtWrapper.GenerateToken(u)

	if err != nil {
		return UserAuth{}, err
	}

	userAuth.Token = signedToken

	return userAuth, nil
}
func (aw *Wrapper) GenerateToken(user entity.User) (signedToken string, err error) {
	claim := &Claim{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
			Issuer:    aw.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err = token.SignedString([]byte(aw.Secret))

	return
}

func (aw *Wrapper) ValidateToken(signedToken string) (claim *Claim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(aw.Secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*Claim)

	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("Token is expired")
		return
	}

	return
}
