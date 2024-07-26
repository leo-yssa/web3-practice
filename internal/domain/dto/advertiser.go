package dto

type AdvertiserCreation struct {
	*Advertiser
	Name string `binding:"required" json:"name"`
}

type Advertiser struct {
	Email  string `binding:"required,email" json:"email"`
	Secret string `binding:"required" json:"secret"`
}
