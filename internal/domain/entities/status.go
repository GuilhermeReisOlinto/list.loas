package entities

import (
	"fmt"

	"github.com/GuilhermeReisOlinto/buscar_emprestimo/internal/infra/entities"
)

type Status interface {
	List(analysis []entities.AnalisysProposal)
}

type EntitiesStatus struct{}

func (es EntitiesStatus) List(analysis []entities.AnalisysProposal) (string, error) {

	var a string
	fmt.Println(analysis)
	for i := 0; i <= len(analysis); i++ {
		statusAnalysis(analysis[i].Status_proposal)
	}

	return a, nil
}

type StatusID int

const (
	Pendente StatusID = iota
	AguardandoAnaliseManual
	EmAnaliseManual
	AguardandoAnaliseSistema
	EmAnaliseSistema
	Aprovado
	Negado
	AprovadoAnaliseManual
	NegadoAnaliseManual
	Questionario
	Expirado
	AguardandoAcessoBio
)

func statusAnalysis(analysis int) string {

	fmt.Println("aqui ", analysis)

	return "aqio"
}
