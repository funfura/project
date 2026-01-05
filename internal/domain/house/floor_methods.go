package house

import (
	"errors"
	"project/internal/domain/common"
)

var ErrWorkNotFound = errors.New("work not found")

func (f *Floor) CompleteWork(systemID, workID common.ID) error {
	for si := range f.Systems {
		if f.Systems[si].ID == systemID {
			for wi := range f.Systems[si].Works {
				if f.Systems[si].Works[wi].ID == workID {
					return f.Systems[si].Works[wi].ComleteSafe()
				}
			}
		}
	}
	return ErrWorkNotFound
}
