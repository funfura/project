package house

import "project/internal/domain/work"

type Apartment struct {
	Number int
	Works  []work.Work
}
