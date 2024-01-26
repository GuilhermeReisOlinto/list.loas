package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/connection/databases"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
)

func GetProposal(client_id int) ([]entities.Proposal, error) {
	conn := databases.Connection()
	if conn == nil {
		log.Println("Error: failed to establish database connection")
		return nil, errors.New("failed to establish database connection")
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT id_proposta, contrato FROM cad_proposta WHERE id_cliente=$1", client_id)
	if err != nil {
		log.Printf("Error in proposal query: %v\n", err)
		return nil, fmt.Errorf("Error in proposal query: %v", err)
	}
	defer rows.Close()

	var proposals []entities.Proposal

	for rows.Next() {
		var proposal entities.Proposal
		var contract sql.NullString

		err := rows.Scan(&proposal.Id_proposal, &contract)
		if err != nil {
			log.Printf("Error scanning row: %v\n", err)
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
		log.Printf("Error iterating over rows: %v\n", err)
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}

	return proposals, nil
}
