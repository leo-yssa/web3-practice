package repository

import (
	"errors"
	"web3-practice/internal/domain/dao"

	"gorm.io/gorm"
)

type AdvertiserRepository interface {
	CreateAdvertiser(advertiser *dao.Advertiser, tx *gorm.DB) error
	FindAdvertiserByEmail(email string) ([]*dao.Advertiser, error)
}

func NewAdvertiserRepository(rdb *gorm.DB) AdvertiserRepository {
	return &advertiserRepository{
		rdb: rdb,
	}
}

type advertiserRepository struct {
	rdb *gorm.DB
}

func (ar *advertiserRepository) CreateAdvertiser(advertiser *dao.Advertiser, tx *gorm.DB) error {
	if advertiser.Id == "" {
		return errors.New("invalid key")
	}
	if tx != nil {
		return tx.Create(advertiser).Error
	}
	return ar.rdb.Create(advertiser).Error
}

func (ar *advertiserRepository) FindAdvertiserByEmail(email string) ([]*dao.Advertiser, error) {
	if email == "" {
		return nil, errors.New("invalid key")
	}
	var advertisers []*dao.Advertiser
	err := ar.rdb.Where("email = ?", email).Find(&advertisers).Error
	return advertisers, err
}
