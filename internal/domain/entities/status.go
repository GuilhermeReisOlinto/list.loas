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

	fmt.Println(analysis)
	if analysis == int(AguardandoAnaliseManual) {
		fmt.Println("aqui ", analysis)
	}

	if analysis == int(EmAnaliseManual) {
		fmt.Println("aqui 1", analysis)
	}

	if analysis == int(AguardandoAnaliseSistema) {
		fmt.Println("aqui 2", analysis)
	}

	if analysis == int(EmAnaliseSistema) {
		fmt.Println("aqui 3", analysis)
	}

	if analysis == int(Aprovado) {
		fmt.Println("aqui 4", analysis)
	}

	if analysis == int(Negado) {
		fmt.Println("aqui 5", analysis)
		//Denied(analysis)
	}

	if analysis == int(AprovadoAnaliseManual) {
		fmt.Println("aqui 6", analysis)
	}

	if analysis == int(NegadoAnaliseManual) {
		fmt.Println("aqui 7", analysis)
	}

	if analysis == int(Questionario) {
		fmt.Println("aqui 8", analysis)
	}

	if analysis == int(Expirado) {
		fmt.Println("aqui 9", analysis)
	}

	if analysis == int(AguardandoAcessoBio) {
		fmt.Println("aqui 10", analysis)
	}

	return "deu bom "
}

// type IDenied struct {
// 	id_proposta              int
// 	status_proposta          string
// 	message                  string
// 	continuar_simulacao      bool
// 	bloquear_novo_emprestimo bool
// }

// func Denied(analysis) {
// 	return IDenied{
// 		// id_proposta: analysis.id_proposta,
// 		// status_proposta: analysis,
// 		// message: 'Infelizmente não foi possível liberar uma proposta no momento!\nNão fique triste, vamos te avisar aqui no APP assim que disponibilizar algum valor para você! 💚',
// 	}
// }
