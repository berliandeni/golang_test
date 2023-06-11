package db

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"

	"github.com/jinzhu/gorm"
)

type pinjamanDb struct {
	db *gorm.DB
}

func NewPinjamanDb(db *gorm.DB) repository.PinjamanRepository {
	return &pinjamanDb{db}
}

func (ur *pinjamanDb) GetById(in int) (*model.PinjamanTrx, error) {

	resp := &model.PinjamanTrx{Id: in}
	err := ur.db.First(resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ur *pinjamanDb) GetLimitByIdKonsumen(in int) (*model.PinjamanLimit, error) {
	resp := &model.PinjamanLimit{}
	err := ur.db.Where("konsumen_id = ?", in).Find(resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ur *pinjamanDb) Create(u *model.PinjamanLimit) (*model.PinjamanLimit, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *pinjamanDb) CreateTrx(u *model.PinjamanTrx) (*model.PinjamanTrx, error) {
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}
