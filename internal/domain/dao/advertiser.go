package dao

import "time"

type Advertiser struct {
	Id      string `gorm:"column:id; primaryKey" json:"id"`
	Email   string `gorm:"column:email; not null" binding:"required,email" json:"email"`
	Secret  string `gorm:"column:secret; not null" binding:"required" json:"secret"`
	Name    string `gorm:"column:name" binding:"required" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	// Campaigns []Campaign `gorm:"foreignKey:AdvertiserId" json:"campaigns,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at; autoCreateTime; <-:create" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; autoCreateTime" json:"updated_at"`
}
