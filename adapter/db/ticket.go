package db

import (
	"errors"
	"golang_test/domain/model"
	"golang_test/domain/repository"

	"github.com/jinzhu/gorm"
)

type ticketDb struct {
	db *gorm.DB
}

func NewTicketDb(db *gorm.DB) repository.TicketRepository {
	return &ticketDb{db}
}

func (ur *ticketDb) GetProductById(in int) (*model.Product, error) {
	resp := &model.Product{Id: in}
	err := ur.db.First(resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ur *ticketDb) CreateTrx(m *model.TicketTrx) (*model.TicketTrx, error) {

	tx := ur.db.Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	pr := &model.Product{
		Id: m.IdProduct,
	}
	if err := tx.First(pr).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if pr.Stock < 1 {
		tx.Rollback()
		return nil, errors.New("stok habis")
	} else {
		pr.Stock = pr.Stock - 1
		if err := tx.Save(&pr).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Create(m).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return m, nil
}
