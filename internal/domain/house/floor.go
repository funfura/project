package house

import (
	"project/internal/domain/house/floorworks"
)

type Floor struct {
	Number     int
	Apartments []Apartment
	Systems    []floorworks.FloorSystem
}
