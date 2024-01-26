package domain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/repository"
	"github.com/go-chi/chi"
)

func ListLoansDomain(w http.ResponseWriter, r *http.Request) {
	cpf := chi.URLParam(r, "cpf")

	respLoans, err := repository.GetLoans(cpf)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if len(respLoans) > 0 {

		id_cliente := respLoans[0].Id_cliente

		resp, _ := getProposalRepository(id_cliente, w)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func getProposalRepository(id_cliente int, w http.ResponseWriter) ([]entities.AnalisysProposal, error) {
	respProposal, err := repository.GetProposal(id_cliente)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	var dataAnalisys []entities.AnalisysProposal

	for i := 0; i < len(respProposal); i++ {
		value := respProposal[i]

		respAnalisys, err := repository.GetAnalisys(value.Id_proposal)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(respAnalisys)

		dataAnalisys = append(dataAnalisys, respAnalisys...)
	}

	return dataAnalisys, nil
}
