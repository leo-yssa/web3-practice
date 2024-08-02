package service_test

import (
	"testing"
	"web3-practice/internal/config"
	"web3-practice/internal/service"

	"github.com/stretchr/testify/assert"
)

func loadTokenService() service.TokenService {
	return service.NewTokenService(&config.Config{
		Jwt: &config.JwtConfig{
			Secret: "secret",
			Access: &config.JwtTokenConfig{
				Duration: "1m",
			},
			Refresh: &config.JwtTokenConfig{
				Duration: "60m",
			},
		},
	})
}
func TestIssue(t *testing.T) {
	ts := loadTokenService()
	_, err := ts.Issue("user1")
	assert.NoError(t, err)
}

func TestVerifyAccessToken(t *testing.T) {
	ts := loadTokenService()
	token, err := ts.Issue("user2")
	assert.NoError(t, err)
	assert.Equal(t, "user2", ts.VerifyAccessToken(token.Access))
}

func TestVerifyRefreshToken(t *testing.T) {
	ts := loadTokenService()
	token, err := ts.Issue("user3")
	assert.NoError(t, err)
	assert.Equal(t, "user3", ts.VerifyRefreshToken(token.Refresh))
}
