package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"web3-practice/internal/config"
	"web3-practice/internal/domain/dao"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

const (
	GOOGLE_SCOPE_EMAIL   = "https://www.googleapis.com/auth/userinfo.email"
	GOOGLE_SCOPE_PROFILE = "https://www.googleapis.com/auth/userinfo.profile"
	GOOGLE_USER_URL      = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="
	YOUTUBE              = "youtube"
	LOGIN                = "login"
)

type GoogleService interface {
	AuthCodeURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	UserInfo(accessToken string) (*dao.Audience, error)
}

type googleService struct {
	cfg *oauth2.Config
}

func NewGoogleService(cfg *config.Config, scope string) GoogleService {
	var redirectURL string
	var scopes = []string{}
	switch scope {
	case LOGIN:
		redirectURL = cfg.Oauth.Google.Redirect.Login
		scopes = append(scopes, GOOGLE_SCOPE_EMAIL, GOOGLE_SCOPE_PROFILE)
	case YOUTUBE:
		redirectURL = cfg.Oauth.Google.Redirect.Youtube
		scopes = append(scopes, youtube.YoutubeReadonlyScope)
	}
	return &googleService{
		cfg: &oauth2.Config{
			ClientID:     cfg.Oauth.Google.Client.Id,
			ClientSecret: cfg.Oauth.Google.Client.Secret,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint:     google.Endpoint,
		},
	}
}

func (gs *googleService) AuthCodeURL(state string) string {
	return gs.cfg.AuthCodeURL(state)
}

func (gs *googleService) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return gs.cfg.Exchange(ctx, code)
}

func (gs *googleService) UserInfo(accessToken string) (*dao.Audience, error) {
	resp, err := http.Get(GOOGLE_USER_URL + accessToken)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	user := &dao.Audience{}
	err = json.Unmarshal(body, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
