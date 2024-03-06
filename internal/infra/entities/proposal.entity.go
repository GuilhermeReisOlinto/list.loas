package entities

type Proposal struct {
	Proposal_id       int    `json:"proposal_id"`
	Contract          string `json:"contract"`
	Inclusion_form_id int    `json:"Inclusion_form_id"`
}
