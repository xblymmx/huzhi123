package utils

import (
	"crypto/sha256"
	"crypto/rand"
	"fmt"
)

func generateRandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateSalt() (string, error) {
	b, err := generateRandBytes(256)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func GenerateEncryptedPassword(pwd string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}

	h := sha256.New()
	saltedPwd := pwd + salt
	h.Write([]byte(saltedPwd))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}


