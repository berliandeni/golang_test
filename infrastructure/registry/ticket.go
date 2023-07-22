package registry

import (
	"golang_test/adapter/controller"
	"golang_test/adapter/db"
	ticketusecase "golang_test/usecase/ticket_usecase"
)

func (r *registry) NewTicketController() controller.Ticket {
	u := ticketusecase.NewTicketUsecase(
		db.NewTicketDb(r.db),
	)

	return controller.NewTicketController(u)
}
