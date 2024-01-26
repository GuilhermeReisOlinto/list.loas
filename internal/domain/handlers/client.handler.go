package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/repository"
	"github.com/go-chi/chi"
)

func HandlerClient(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")

	clients, err := repository.GetClient(cpf)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(clients) <= 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	client_id := clients[0].Client_id

	analysisData, _ := HandlerProposals(client_id, w)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysisData)
}
