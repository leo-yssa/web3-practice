package hash_test

import (
	"testing"
	"web3-practice/pkg/hash"

	"github.com/stretchr/testify/assert"
)

func TestHmac(t *testing.T) {
	t.Run("src data `Leo Ko`s base64 string should be `l0dIrsTnolq2bqu8fL6ld5I4+bX9G/dysxAyByFvcEk=`", func(t *testing.T) {
		hmacService := hash.NewHmacService(hash.NewHasher(hash.SHA256).Hasher(), []byte("secret"))
		hash, err := hmacService.Hash([]byte("Leo Ko"))
		assert.NoError(t, err)
		expected := "l0dIrsTnolq2bqu8fL6ld5I4+bX9G/dysxAyByFvcEk="
		assert.EqualValues(t, expected, hash.ToBase64())
	})
}
