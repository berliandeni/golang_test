package repository

import "golang_test/domain/model"

type PinjamanRepository interface {
	GetById(int) (*model.PinjamanTrx, error)
	GetLimitByIdKonsumen(int) (*model.PinjamanLimit, error)
	Create(*model.PinjamanLimit) (*model.PinjamanLimit, error)
	CreateTrx(*model.PinjamanTrx) (*model.PinjamanTrx, error)
}
