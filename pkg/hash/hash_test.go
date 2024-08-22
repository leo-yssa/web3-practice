package hash_test

import (
	"testing"
	"web3-practice/pkg/hash"

	"github.com/stretchr/testify/assert"
)

func TestHasher(t *testing.T) {
	t.Run("should be no error", func(t *testing.T) {
		hasher := hash.NewHasher(hash.SHA256)
		_, err := hasher.Hash([]byte("Leo Ko"))
		assert.NoError(t, err)
	})
	t.Run("src data `Leo Ko`s hex string should be equal `c99aeaa34306ed7c999c29b517af4637d2e50886c09a1adffd6108f34c2917eb`", func(t *testing.T) {
		hasher := hash.NewHasher(hash.SHA256)
		hash, _ := hasher.Hash([]byte("Leo Ko"))
		expected := "c99aeaa34306ed7c999c29b517af4637d2e50886c09a1adffd6108f34c2917eb"
		assert.EqualValues(t, expected, hash.ToHex())
	})
	t.Run("src data `Leo Ko`s base64 string should be equal `yZrqo0MG7XyZnCm1F69GN9LlCIbAmhrf/WEI80wpF+s=`", func(t *testing.T) {
		hasher := hash.NewHasher(hash.SHA256)
		hash, _ := hasher.Hash([]byte("Leo Ko"))
		expected := "yZrqo0MG7XyZnCm1F69GN9LlCIbAmhrf/WEI80wpF+s="
		assert.EqualValues(t, expected, hash.ToBase64())
	})
	t.Run("src data `Leo Ko`s length should be equal 32", func(t *testing.T) {
		hasher := hash.NewHasher(hash.SHA256)
		hash, _ := hasher.Hash([]byte("Leo Ko"))
		expected := 32
		assert.EqualValues(t, expected, len(hash.ToBytes()))
	})
}
