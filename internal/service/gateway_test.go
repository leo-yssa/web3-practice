package service_test

import (
	"testing"
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/service"

	"github.com/stretchr/testify/assert"
)

func loadFixtures() (s service.GatewayService) {
	return service.NewGatewayService(
		"http://15.164.146.25:8080",
	)
}

func login(s service.GatewayService, loginRequest *dto.LoginRequest) (*dto.LoginResponse, error) {
	return s.Login(&dto.LoginRequest{
		UserType:                   "ADVERTISER",
		Email:                      "test1@test.com",
		Name:                       "test",
		SnsProvider:                "LOCAL",
		Image:                      "",
		ExternalId:                 "01J31TB27Q042G8WBC5E6QJ8X7",
		SocialProvidedAccessToken:  "",
		SocialProvidedRefreshToken: "",
	})
}
func TestLogin(t *testing.T) {
	s := loadFixtures()
	r, err := login(s, &dto.LoginRequest{
		UserType:                   "ADVERTISER",
		Email:                      "test1@test.com",
		Name:                       "test",
		SnsProvider:                "LOCAL",
		Image:                      "",
		ExternalId:                 "01J31TB27Q042G8WBC5E6QJ8X7",
		SocialProvidedAccessToken:  "",
		SocialProvidedRefreshToken: "",
	})
	t.Log(r)
	assert.NoError(t, err)
}

func TestMintDHN(t *testing.T) {
	s := loadFixtures()
	lr, err := login(s, &dto.LoginRequest{
		UserType:                   "USER",
		Email:                      "test1@google.com",
		Name:                       "test",
		SnsProvider:                "GOOGLE",
		Image:                      "",
		ExternalId:                 "01J31TB27Q042G8WBC5E6QJ8X8",
		SocialProvidedAccessToken:  "adsfasdfasdfasdfasdf",
		SocialProvidedRefreshToken: "asdfasdfasdfasdfasf",
	})
	t.Log(lr)
	assert.NoError(t, err)
	mr, err := s.MintDHN(&dto.MintDHNRequest{
		Metadata: &dto.Metadata{
			Image:      "",
			Attributes: []string{},
		},
		MetadataUri: "",
		ImageUri:    "",
		SegmentId:   "01J31TB27Q042G8WBC5E6QJ8X8",
	}, lr.AccessToken)
	t.Log(mr)
	assert.NoError(t, err)
}
