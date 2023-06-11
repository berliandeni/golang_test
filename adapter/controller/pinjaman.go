package controller

import (
	"golang_test/domain/model"
	pinjamanusecase "golang_test/usecase/pinjaman_usecase"
	"net/http"
	"strconv"
)

type pinjamanController struct {
	pinjamanUsecase pinjamanusecase.Pinjaman
}

type Pinjaman interface {
	GetById(Context) error
	CreateTrx(Context) error
}

func NewPinjamanController(us pinjamanusecase.Pinjaman) Pinjaman {
	return &pinjamanController{us}
}

func (p *pinjamanController) GetById(ctx Context) error {
	in, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	resp, err := p.pinjamanUsecase.GetById(in)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (p *pinjamanController) CreateTrx(ctx Context) error {
	var params model.PinjamanTrxReq

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	resp, err := p.pinjamanUsecase.CreateTrx(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, resp)
}
