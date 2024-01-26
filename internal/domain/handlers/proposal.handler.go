package handlers

import (
	"net/http"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/repository"
)

func HandlerProposals(id_cliente int, w http.ResponseWriter) ([]entities.AnalisysProposal, error) {
	proposals, err := repository.GetProposals(id_cliente)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil, err
	}

	var analysisData []entities.AnalisysProposal

	for i := 0; i < len(proposals); i++ {
		value := proposals[i]

		analysis, err := repository.GetAnalisys(value.Proposal_id)
		if err != nil {
			continue
		}

		analysisData = append(analysisData, analysis...)
	}

	return analysisData, nil
}
