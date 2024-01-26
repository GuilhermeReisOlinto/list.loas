package presentation

import (
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/domain/handlers"
	"github.com/go-chi/chi"
)

func PresentationList(r chi.Router) {

	r.Get("/list/loans/{cpf}", handlers.HandlerClient)
}
