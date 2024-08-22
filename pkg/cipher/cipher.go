package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

type CipherTextError string

func (t CipherTextError) Error() string {
	return "crypto/aes: invalid cipher text" + string(t)
}

type Cipher interface {
	CBCEncrypt(plainText string) (string, error)
	CBCDecrypt(cipherText string) ([]byte, error)
}

func NewCipher(key, iv string) (Cipher, error) {
	switch len(key) {
	case 16, 24, 32:
		block, err := aes.NewCipher([]byte(key))
		if err != nil {
			return nil, err
		}
		return &aesCipher{
			iv:    []byte(iv),
			block: block,
		}, nil
	}
	return nil, aes.KeySizeError(len(key))
}

type aesCipher struct {
	iv    []byte
	block cipher.Block
}

func (ac *aesCipher) CBCEncrypt(plainText string) (string, error) {
	encrypter := cipher.NewCBCEncrypter(ac.block, ac.iv)
	paddedPlainText := ac.padPKCS7([]byte(plainText), encrypter.BlockSize())
	cipherText := make([]byte, len(paddedPlainText))
	encrypter.CryptBlocks(cipherText, paddedPlainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (ac *aesCipher) padPKCS7(plain []byte, blockSize int) []byte {
	padding := blockSize - len(plain)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plain, padText...)
}

func (ac *aesCipher) CBCDecrypt(cipherText string) ([]byte, error) {
	if strings.TrimSpace(cipherText) == "" {
		return nil, CipherTextError(cipherText)
	}
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return nil, err
	}
	decrypter := cipher.NewCBCDecrypter(ac.block, ac.iv)
	plainText := make([]byte, len(decodedCipherText))
	decrypter.CryptBlocks(plainText, decodedCipherText)
	return ac.trimPKCS5(plainText), nil
}

func (ac *aesCipher) trimPKCS5(plainText []byte) []byte {
	padding := plainText[len(plainText)-1]
	return plainText[:len(plainText)-int(padding)]
}
