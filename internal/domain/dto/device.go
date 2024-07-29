package dto

type Device struct {
	Id         string `binding:"required" json:"id"`
	Name       string `binding:"required" json:"name"`
	Kind       string `binding:"required" json:"kind"`
	Size       string `binding:"required" json:"size"`
	Resolution string `binding:"required" json:"resolution"`
	Output     string `binding:"required" json:"output"`
	Channel    string `binding:"required" json:"channel"`
}
