package dto

type Jwt struct {
	Access  string `json:"assess_token"`
	Refresh string `json:"refresh_token"`
}
