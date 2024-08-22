package hash

import (
	"crypto/hmac"
	"encoding/base64"
	"hash"
)

type HmacService interface {
	Hash(src []byte) (HmacService, error)
	ToBase64() string
}

func NewHmacService(hasher func() hash.Hash, secret []byte) HmacService {
	return &hmacService{
		hasher: hmac.New(hasher, secret),
	}
}

type hmacService struct {
	hasher hash.Hash
}

func (hs *hmacService) Hash(src []byte) (HmacService, error) {
	hs.hasher.Reset()
	if _, err := hs.hasher.Write(src); err != nil {
		return nil, err
	}
	return hs, nil
}

func (hs *hmacService) ToBase64() string {
	return base64.StdEncoding.EncodeToString(hs.hasher.Sum(nil))
}
