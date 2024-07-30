package service

import (
	"time"
	"web3-practice/internal/config"
	"web3-practice/internal/domain/dto"
	"web3-practice/pkg/util"

	"github.com/golang-jwt/jwt/v4"
)

const (
	access  = "access"
	refresh = "refresh"
)

type TokenService interface {
	Issue(aud string) (*dto.Jwt, error)
	VerifyAccessToken(tokenString string) string
	VerifyRefreshToken(tokenString string) string
}

type jwtService struct {
	secret          []byte
	accessDuration  string
	refreshDuration string
}

func NewTokenService(cfg *config.Config) TokenService {
	return &jwtService{
		secret:          []byte(cfg.Jwt.Secret),
		accessDuration:  cfg.Jwt.Access.Duration,
		refreshDuration: cfg.Jwt.Refresh.Duration,
	}
}

func (j *jwtService) makeToken(sub, aud string) (string, error) {
	mapClaims := jwt.MapClaims{}
	mapClaims["sub"] = sub
	mapClaims["iat"] = util.Unix()
	mapClaims["aud"] = aud
	var ds string
	if sub == access {
		ds = j.accessDuration
	} else {
		ds = j.refreshDuration
	}
	d, err := time.ParseDuration(ds)
	if err != nil {
		return "", err
	}
	mapClaims["exp"] = time.Now().Add(d).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, mapClaims)
	return token.SignedString(j.secret)
}

func (j *jwtService) Issue(aud string) (*dto.Jwt, error) {
	var err error
	token := &dto.Jwt{}
	if token.Access, err = j.makeToken(access, aud); err != nil {
		return nil, err
	}
	if token.Refresh, err = j.makeToken(refresh, aud); err != nil {
		return nil, err
	}
	return token, err
}

func (j *jwtService) VerifyAccessToken(tokenString string) string {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	sub, _ := claims["sub"].(string)
	if err != nil || token == nil || !token.Valid || sub != access {
		return ""
	}
	aud, _ := claims["aud"].(string)
	return aud
}

func (j *jwtService) VerifyRefreshToken(tokenString string) string {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	sub, _ := claims["sub"].(string)
	if err != nil || token == nil || !token.Valid || sub != refresh {
		return ""
	}
	aud, _ := claims["aud"].(string)
	return aud
}
