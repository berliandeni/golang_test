package konsumenusecase

import (
	"golang_test/domain/model"
	"golang_test/domain/repository"
	"strconv"
)

type konsumenUsecase struct {
	konsumenRepository repository.KonsumenRepository
	pinjamanRepository repository.PinjamanRepository
}

type Konsumen interface {
	GetById(u int) (*model.KonsumenResp, error)
	GetAll(u []*model.Konsumen) ([]*model.Konsumen, error)
	Create(u *model.Konsumen) (*model.KonsumenResp, error)
}

func NewKonsumenUsecase(r repository.KonsumenRepository, p repository.PinjamanRepository) Konsumen {
	return &konsumenUsecase{r, p}
}

func (uk *konsumenUsecase) GetById(in int) (*model.KonsumenResp, error) {
	k, err := uk.konsumenRepository.GetById(in)
	if err != nil {
		return nil, err
	}

	pl, err := uk.pinjamanRepository.GetLimitByIdKonsumen(in)
	if err != nil {
		return nil, err
	}

	resp := &model.KonsumenResp{
		Konsumen:      *k,
		PinjamanLimit: *pl,
	}

	return resp, nil
}

func (uk *konsumenUsecase) GetAll(in []*model.Konsumen) ([]*model.Konsumen, error) {
	in, err := uk.konsumenRepository.GetAll(in)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func (uk *konsumenUsecase) Create(k *model.Konsumen) (*model.KonsumenResp, error) {
	k, err := uk.konsumenRepository.Create(k)
	if err != nil {
		return nil, err
	}

	gaji, err := strconv.ParseFloat(k.Gaji, 64)
	if err != nil {
		return nil, err
	}

	pl := &model.PinjamanLimit{
		KonsumenId: k.Id,
		Tenor1:     gaji * 0.1,
		Tenor2:     gaji * 0.2,
		Tenor3:     gaji * 0.4,
		Tenor4:     gaji * 0.6,
	}

	pl, err = uk.pinjamanRepository.Create(pl)
	if err != nil {
		return nil, err
	}

	resp := &model.KonsumenResp{
		Konsumen:      *k,
		PinjamanLimit: *pl,
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
