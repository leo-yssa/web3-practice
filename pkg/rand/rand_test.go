package rand_test

import (
	"testing"
	"web3-practice/pkg/rand"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	_, err := rand.MakeState()
	assert.NoError(t, err)
	_, err = rand.MakeUUID()
	assert.NoError(t, err)
}
