package dto

type GoogleAuthState struct {
	Uuid  string `json:"uuid"`
	State string `json:"state"`
}

type GoogleAuthCodeURL struct {
	*GoogleAuthState
	Url string `json:"url"`
}

type GoogleLogin struct {
	*GoogleAuthState
	*Device
	Code string `json:"code"`
}
