package repository

import "golang_test/domain/model"

type TicketRepository interface {
	GetProductById(int) (*model.Product, error)
	// CreateProduct(*model.Product) (*model.Product, error)
	CreateTrx(*model.TicketTrx) (*model.TicketTrx, error)
}
