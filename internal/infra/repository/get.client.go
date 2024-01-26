package repository

import (
	"errors"
	"fmt"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/connection/databases"
	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
)

func GetClient(cpf string) (clientData []entities.Client, err error) {
	conn := databases.Connection()
	defer conn.Close()

	if conn == nil {
		return nil, errors.New("failed to establish database connection")
	}

	rows, err := conn.Query("SELECT Id_cliente, Nome FROM cad_cliente WHERE cpf=$1", cpf)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entities.Client
		if err := rows.Scan(&client.Client_id, &client.Name); err != nil {
			continue
		}
		clientData = append(clientData, client)
	}

	return clientData, nil
}
