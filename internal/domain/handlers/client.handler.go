package handlers

import (
	"fmt"
	"net/http"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/domain/entities"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/repository"
	"github.com/go-chi/chi"
)

type Client interface {
	Client(w http.ResponseWriter, r *http.Request)
}

type HandlerClient struct{}

func (hc HandlerClient) Client(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")

	clientRepository := repository.Client{}

	clients, err := clientRepository.Get(cpf)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(clients) <= 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	client_id := clients[0].Client_id

	handlerProposals := HandlerProposals{}
	analysisData, _ := handlerProposals.Proposals(client_id, w)

	entitiesStatus := entities.EntitiesStatus{}
	status, _ := entitiesStatus.List(analysisData)
	fmt.Println(status)
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(status)
}
