package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func Sha256ToBase64(data []byte) (string, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}

func HmacToBase64(key []byte, data []byte) (string, error) {
	h := hmac.New(sha256.New, key)
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func Sha256ToHex(data []byte) (string, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
