package schemas

import "gorm.io/gorm"

type Outcome struct {
	gorm.Model

	Name string
	Amount int64
}