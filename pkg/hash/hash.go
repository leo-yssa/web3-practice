package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

type HasherType int

const (
	SHA256 = iota
)

type Hasher interface {
	Hash(src []byte) (Hasher, error)
	Hasher() func() hash.Hash
	ToBase64() string
	ToBytes() []byte
	ToHex() string
}

func NewHasher(hasherType HasherType) Hasher {
	switch hasherType {
	case SHA256:
		return &sha256Hasher{
			hasher: sha256.New(),
		}
	}
	return nil
}

type sha256Hasher struct {
	hasher hash.Hash
}

func (sh *sha256Hasher) Hasher() func() hash.Hash {
	return func() hash.Hash {
		return sha256.New()
	}
}

func (sh *sha256Hasher) Hash(src []byte) (Hasher, error) {
	sh.hasher.Reset()
	if _, err := sh.hasher.Write(src); err != nil {
		return nil, err
	}
	return sh, nil
}

func (sh *sha256Hasher) ToBytes() []byte {
	return sh.hasher.Sum(nil)
}

func (sh *sha256Hasher) ToHex() string {
	return hex.EncodeToString(sh.hasher.Sum(nil))
}

func (sh *sha256Hasher) ToBase64() string {
	return base64.StdEncoding.EncodeToString(sh.hasher.Sum(nil))
}
