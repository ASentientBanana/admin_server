package util

import (
	"errors"
	"os"
)

type UserEnv struct {
	Username string
	Password string
	Secret   string
}

func LoadUserEnv() (UserEnv, error) {

	admin_user := os.Getenv("ADMIN_USER")
	admin_password := os.Getenv("ADMIN_PASSWORD")
	secret := os.Getenv("ADMIN_SECRET")

	if admin_password == "" || admin_user == "" || secret == "" {
		return UserEnv{}, errors.New("Problem loading environment variables.")
	}

	return UserEnv{Username: admin_user, Password: admin_password, Secret: secret}, nil

}
