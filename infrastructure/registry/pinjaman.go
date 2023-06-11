package registry

import (
	"golang_test/adapter/controller"
	"golang_test/adapter/db"
	pinjamanusecase "golang_test/usecase/pinjaman_usecase"
)

func (r *registry) NewPinjamanController() controller.Pinjaman {
	u := pinjamanusecase.NewPinjamanUsecase(
		db.NewPinjamanDb(r.db),
	)

	return controller.NewPinjamanController(u)
}
