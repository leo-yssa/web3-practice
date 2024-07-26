package util

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(secret string) (string, error) {
	s, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func ComparePassword(hash, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
}
