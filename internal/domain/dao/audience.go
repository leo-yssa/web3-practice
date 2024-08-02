package dao

import "time"

type Audience struct {
	Id               string `gorm:"column:id; primaryKey" json:"id"`
	Email            string `gorm:"column:email; not null" binding:"required,email" json:"email"`
	Name             string `gorm:"column:name; not null" binding:"required" json:"name"`
	GivenName        string `gorm:"column:given_name" json:"given_name"`
	Picture          string `gorm:"column:picture" json:"picture"`
	Subject          string `gorm:"column:subject" json:"subject"`
	BirthDate        string `gorm:"column:birth_dt" json:"birth_dt"`
	Gender           string `gorm:"column:gender" json:"gender"`
	Score            uint64 `gorm:"column:score;default:0" json:"score"`
	Face             string `gorm:"column:face" json:"face"`
	Clothes          string `gorm:"column:clothes" json:"clothes"`
	Background       string `gorm:"column:background" json:"background"`
	PointBalance     uint64 `gorm:"column:point_balance" json:"point_balance"`
	MarketingConsent bool   `gorm:"column:marketing_consent" json:"marketing_consent"`
	// AudienceLinked   []AudienceLinked `gorm:"foreignKey:AudienceId" json:"linked"`
	// RootInventory    RootInventory    `gorm:"foreignKey:AudienceId" json:"root_inventory"`
	CreatedAt time.Time `gorm:"column:created_at; autoCreateTime; <-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; autoCreateTime" json:"updated_at"`
}
