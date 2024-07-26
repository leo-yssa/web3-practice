package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GetState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
