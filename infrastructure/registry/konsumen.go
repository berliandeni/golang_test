package registry

import (
	"golang_test/adapter/controller"
	"golang_test/adapter/db"
	konsumenusecase "golang_test/usecase/konsumen_usecase"
)

func (r *registry) NewKonsumenController() controller.Konsumen {
	u := konsumenusecase.NewKonsumenUsecase(
		db.NewkonsumenDb(r.db),
		db.NewPinjamanDb(r.db),
	)

	return controller.NewKonsumenController(u)
}
