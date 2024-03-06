package repository

import (
	"errors"
	"fmt"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/connection/databases"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
)

type IAnalysis interface {
	Get(proposal_id int) ([]entities.AnalisysProposal, error)
}

type Analysis struct{}

func (a Analysis) Get(proposal_id int) ([]entities.AnalisysProposal, error) {
	conn := databases.Connection()

	if conn == nil {
		return nil, errors.New("failed to establish database connection")
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT id_analise_proposta, id_status_analise_proposta, id_proposta FROM cad_analise_proposta WHERE id_proposta=$1", proposal_id)
	if err != nil {
		return nil, fmt.Errorf("Error in analysis query: %v", err)
	}

	defer rows.Close()

	var analysies []entities.AnalisysProposal

	for rows.Next() {
		var analysis entities.AnalisysProposal

		if err = rows.Scan(&analysis.Analisys_proposal_id, &analysis.Status_proposal, &analysis.Proposal_id); err != nil {
			continue
		}

		analysies = append(analysies, analysis)
	}

	return analysies, nil
}
