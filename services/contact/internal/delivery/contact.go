package delivery

import (
	"architecture_go/services/contact/internal/repository/controller"
	"architecture_go/services/contact/internal/repository/storage/postgres"
	use_case "architecture_go/services/contact/internal/use-case"
)

func (r *registry) NewContactController() controller.Contact {
	c := use_case.NewContactUseCase(
		postgres.NewContactRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewContactsController(c)
}
