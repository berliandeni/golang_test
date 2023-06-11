package db

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"

	"github.com/jinzhu/gorm"
)

type konsumenDb struct {
	db *gorm.DB
}

func NewkonsumenDb(db *gorm.DB) repository.KonsumenRepository {
	return &konsumenDb{db}
}

func (ur *konsumenDb) GetById(in int) (*model.Konsumen, error) {

	resp := &model.Konsumen{Id: in}
	err := ur.db.First(resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ur *konsumenDb) GetAll(u []*model.Konsumen) ([]*model.Konsumen, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *konsumenDb) Create(u *model.Konsumen) (*model.Konsumen, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
