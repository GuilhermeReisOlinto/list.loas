package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/connection/databases"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
)

type IProposals interface {
	Get(client_id int) ([]entities.Proposal, error)
}

type Proposal struct{}

func (p Proposal) Get(client_id int) ([]entities.Proposal, error) {
	conn := databases.Connection()

	if conn == nil {
		return nil, errors.New("failed to establish database connection")
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id_proposta, contrato, id_forma_inclusao FROM cad_proposta WHERE id_cliente=$1", client_id)
	if err != nil {
		return nil, fmt.Errorf("Error in proposal query: %v", err)
	}
	defer rows.Close()

	var proposals []entities.Proposal

	for rows.Next() {
		var proposal entities.Proposal
		var contract sql.NullString

		err := rows.Scan(&proposal.Proposal_id, &contract, &proposal.Inclusion_form_id)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row: %v", err)
		}

		if contract.Valid {
			proposal.Contract = contract.String
		} else {
			proposal.Contract = ""
		}

		proposals = append(proposals, proposal)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}

	return proposals, nil
}
