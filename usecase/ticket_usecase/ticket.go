package ticketusecase

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"
)

type ticketUsecase struct {
	ticketRepository repository.TicketRepository
}

type Ticket interface {
	GetProductById(in int) (*model.Product, error)
	CreateTrx(u *model.TicketTrx) (*model.TicketTrx, error)
}

func NewTicketUsecase(p repository.TicketRepository) Ticket {
	return &ticketUsecase{p}
}

func (uk *ticketUsecase) GetProductById(in int) (*model.Product, error) {
	p, err := uk.ticketRepository.GetProductById(in)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (up *ticketUsecase) CreateTrx(p *model.TicketTrx) (*model.TicketTrx, error) {

	resp, err := up.ticketRepository.CreateTrx(p)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
