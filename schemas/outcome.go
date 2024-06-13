package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Outcome struct {
	gorm.Model

	Name string
	Amount int64
}

type OutcomeResponse struct {
	ID uint `json:"id"`
	Name uint `json:"name"`
	Amount uint `json:"amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}