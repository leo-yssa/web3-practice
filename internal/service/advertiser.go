package service

import (
	"web3-practice/internal/domain/dao"
	"web3-practice/internal/domain/dto"
	"web3-practice/internal/repository"
	"web3-practice/pkg/util"

	"gorm.io/gorm"
)

type AdvertiserService interface {
	CreateAdvertiser(advertiser *dto.AdvertiserCreation, tx *gorm.DB) error
	FindAdvertiserByEmail(email string) ([]*dao.Advertiser, error)
}

func NewAdvertiserService(repo repository.Repository) AdvertiserService {
	return &advertiserService{
		repo: repo,
	}
}

type advertiserService struct {
	repo repository.Repository
}

func (as *advertiserService) FindAdvertiserByEmail(email string) ([]*dao.Advertiser, error) {
	return as.repo.FindAdvertiserByEmail(email)
}

func (as *advertiserService) CreateAdvertiser(advertiser *dto.AdvertiserCreation, tx *gorm.DB) error {
	secret, err := util.GenerateFromPassword(advertiser.Secret)
	if err != nil {
		return err
	}
	return as.repo.CreateAdvertiser(&dao.Advertiser{
		Id:     util.GenerateULID("AD"),
		Email:  advertiser.Email,
		Secret: secret,
		Name:   advertiser.Name,
	}, tx)
}
