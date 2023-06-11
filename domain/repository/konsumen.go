package repository

import "golang_test/domain/model"

type KonsumenRepository interface {
	GetById(int) (*model.Konsumen, error)
	GetAll([]*model.Konsumen) ([]*model.Konsumen, error)
	Create(*model.Konsumen) (*model.Konsumen, error)
}
