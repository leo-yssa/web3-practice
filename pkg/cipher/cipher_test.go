package cipher_test

import (
	"testing"
	"web3-practice/pkg/cipher"

	"github.com/stretchr/testify/assert"
)

func TestCipter(t *testing.T) {
	source := "A1B2C3D4E5F6G7H8I9J0K1L2M3N4O5P6Q7R8S9T0U1V2W3X4Y5Z6="
	t.Run("invalid key size", func(t *testing.T) {
		_, err := cipher.NewCipher(source[:15], source[len(source)-16:])
		assert.Error(t, err)
	})
	t.Run("valid key size", func(t *testing.T) {
		_, err := cipher.NewCipher(source[:16], source[len(source)-16:])
		assert.NoError(t, err)
	})
	t.Run("plain text `Leo Ko` should be equal `SvwyNh5MbbRaVedxsXLuQQ==`", func(t *testing.T) {
		c, _ := cipher.NewCipher(source[:16], source[len(source)-16:])
		cipherText, err := c.CBCEncrypt("Leo Ko")
		assert.NoError(t, err)
		expected := "SvwyNh5MbbRaVedxsXLuQQ=="
		assert.EqualValues(t, expected, cipherText)
	})
	t.Run("decrypted text `SvwyNh5MbbRaVedxsXLuQQ==` should be equal `Leo Ko`", func(t *testing.T) {
		cipher, _ := cipher.NewCipher(source[:16], source[len(source)-16:])
		plainText, err := cipher.CBCDecrypt("SvwyNh5MbbRaVedxsXLuQQ==")
		assert.NoError(t, err)
		expected := "Leo Ko"
		assert.EqualValues(t, expected, plainText)
	})
}
