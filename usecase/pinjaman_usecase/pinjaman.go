package pinjamanusecase

import (
	"errors"
	"fmt"
	"golang_test/domain/model"
	"golang_test/domain/repository"
)

type pinjamanUsecase struct {
	pinjamanRepository repository.PinjamanRepository
}

type Pinjaman interface {
	GetById(in int) (*model.PinjamanTrx, error)
	CreateTrx(u *model.PinjamanTrxReq) (*model.PinjamanTrx, error)
}

func NewPinjamanUsecase(p repository.PinjamanRepository) Pinjaman {
	return &pinjamanUsecase{p}
}

func (uk *pinjamanUsecase) GetById(in int) (*model.PinjamanTrx, error) {
	k, err := uk.pinjamanRepository.GetById(in)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (up *pinjamanUsecase) CreateTrx(p *model.PinjamanTrxReq) (*model.PinjamanTrx, error) {
	dataLimit, err := up.pinjamanRepository.GetLimitByIdKonsumen(p.KonsumenId)
	if err != nil {
		return nil, err
	}

	bunga := p.PinjamanBody.OTR * 0.1
	admin := p.PinjamanBody.OTR * 0.01
	cicilan := (p.PinjamanBody.OTR + bunga + admin) / float64(p.Tenor)
	var limit float64
	if p.Tenor == 1 {
		limit = dataLimit.Tenor1
	} else if p.Tenor == 2 {
		limit = dataLimit.Tenor2
	} else if p.Tenor == 3 {
		limit = dataLimit.Tenor3
	} else {
		limit = dataLimit.Tenor4
	}

	fmt.Println(float64(p.Tenor))

	if cicilan > limit {
		return nil, errors.New("transaksi melebihi limit")
	}

	req := &model.PinjamanTrx{
		NoKontak:      p.PinjamanBody.NoKontak,
		OTR:           p.PinjamanBody.OTR,
		NamaAsset:     p.PinjamanBody.NamaAsset,
		AdminFee:      admin,
		JumlahBunga:   bunga,
		JumlahCicilan: cicilan,
	}
	resp, err := up.pinjamanRepository.CreateTrx(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
