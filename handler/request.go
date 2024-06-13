package handler

import "fmt"

func errRequiredParam(p, t string) error {
	return fmt.Errorf("param '%s' of type '%s' is required", p, t)
}

type CreateOutcomeRequest struct {
	Name string `json:"name"`
	Amount int64 `json:"amount"`
}

func (r *CreateOutcomeRequest) Validate() error {
	if r.Name == "nil" && r.Amount <= 0 {
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