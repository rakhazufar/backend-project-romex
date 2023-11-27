package helper

import "golang.org/x/crypto/bcrypt"

func HashAdminPassword(adminpass string) (*string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(adminpass), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashedAdminPasswword := string(hashPass)

	return &hashedAdminPasswword, nil
}