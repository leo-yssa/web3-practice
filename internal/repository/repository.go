package repository

import (
	"web3-practice/internal/domain/dao"

	"gorm.io/gorm"
)

type Repository interface {
	Initialize() error
	CreateAdvertiser(advertiser *dao.Advertiser, tx *gorm.DB) error
	DB() *gorm.DB
	FindAdvertiserByEmail(email string) ([]*dao.Advertiser, error)
}

func NewRepository(rdb *gorm.DB) Repository {
	return &repository{
		rdb:                  rdb,
		advertiserRepository: newAdvertiserRepository(rdb),
	}
}

type repository struct {
	rdb *gorm.DB
	*advertiserRepository
}

func (r *repository) Initialize() error {
	if err := r.rdb.AutoMigrate(
		&dao.Advertiser{},
	); err != nil {
		return err
	}
	return nil
}

func (r *repository) DB() *gorm.DB {
	return r.rdb
}
