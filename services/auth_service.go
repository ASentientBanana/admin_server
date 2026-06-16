package services

import (
	"errors"
	"time"

	"github.com/AsentientBanana/admin/dto"
	"github.com/AsentientBanana/admin/util"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateUser(authDto dto.AuthDto, userEnv util.UserEnv) bool {
	if authDto.Password != userEnv.Password {
		return false
	}

	if authDto.Username != userEnv.Username {
		return false
	}

	return true
}

func AuthenticateUser(authDto dto.AuthDto) (string, error) {

	user_env, err := util.LoadUserEnv()

	if err != nil {

		return "", err
	}

	isUserValid := ValidateUser(authDto, user_env)

	if !isUserValid {
		return "", errors.New("Invalid user")
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_env.Username,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the admin credentials
	return token.SignedString([]byte(user_env.Secret))
}
