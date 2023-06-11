package controller

import (
	"golang_test/domain/model"
	konsumenusecase "golang_test/usecase/konsumen_usecase"
	"net/http"
	"strconv"
)

type konsumenController struct {
	konsumenUsecase konsumenusecase.Konsumen
}

type Konsumen interface {
	GetById(Context) error
	GetAll(Context) error
	Create(Context) error
}

func NewKonsumenController(us konsumenusecase.Konsumen) Konsumen {
	return &konsumenController{us}
}

func (kc *konsumenController) GetById(ctx Context) error {
	p, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return err
	}

	resp, err := kc.konsumenUsecase.GetById(p)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (kc *konsumenController) GetAll(ctx Context) error {
	var resp []*model.Konsumen

	resp, err := kc.konsumenUsecase.GetAll(resp)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (kc *konsumenController) Create(ctx Context) error {
	var params model.Konsumen

	if err := ctx.Bind(&params); err != nil {
		return err
	}

	resp, err := kc.konsumenUsecase.Create(&params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, resp)
}
