package house

import "project/internal/domain/basement"

type House struct {
	ID       string
	Entrance []Entrance
	Basement basement.Basement
}
