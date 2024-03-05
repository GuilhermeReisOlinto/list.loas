package controller

import (
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/domain/handlers"

	"github.com/go-chi/chi"
)

func PresentationList(r chi.Router) {
	handlerClient := handlers.HandlerClient{}
	r.Get("/list/loans/{cpf}", handlerClient.Client)
}
