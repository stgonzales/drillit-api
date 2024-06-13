package handler

import "fmt"

func errRequiredParam(p, t string) error {
	return fmt.Errorf("param '%s' of type '%s' is required", p, t)
}

type OutcomeRequest struct {
	Name string `json:"name"`
	Amount int64 `json:"amount"`
}

func (r *OutcomeRequest) ValidateCreateOutcomeRequest() error {
	if r.Name == "" && r.Amount <= 0 {
		return fmt.Errorf("malformed request body")
	}

	if r.Name == "" {
		return errRequiredParam("name", "string")
	}

	if r.Amount <= 0 {
		return errRequiredParam("amount", "int64")
	}

	return nil
}

func (r *OutcomeRequest) ValidateUpdateOutcomeRequest() error {
	if r.Name != "" || r.Amount > 0 {
		return nil
	}
	 
	return fmt.Errorf("malformed request body, at least one field of type Outcome should be provided")

}