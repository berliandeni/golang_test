package controller

import (
	"golang_test/domain/model"
	ticketusecase "golang_test/usecase/ticket_usecase"
	"net/http"
	"strconv"
)

type ticketController struct {
	ticketUsecase ticketusecase.Ticket
}

type Ticket interface {
	GetProductById(Context) (interface{}, error)
	CreateTrx(Context) error
}

func NewTicketController(us ticketusecase.Ticket) Ticket {
	return &ticketController{us}
}

func (p *ticketController) GetProductById(ctx Context) (interface{}, error) {
	in, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	resp, err := p.ticketUsecase.GetProductById(in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *ticketController) CreateTrx(ctx Context) error {
	var params model.TicketTrx

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	resp, err := p.ticketUsecase.CreateTrx(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, resp)
}
