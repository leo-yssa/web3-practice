package dto

type GoogleAuthCodeURL struct {
	Uuid  string `json:"uuid"`
	State string `json:"state"`
}
type GoogleLogin struct {
	*GoogleAuthCodeURL
	*Device
	Code string `json:"code"`
}
