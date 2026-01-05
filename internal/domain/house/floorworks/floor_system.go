package floorworks

import (
	"project/internal/domain/common"
	"project/internal/domain/work"
)

type FloorSystem struct {
	ID    common.ID
	Name  string
	Works []work.Work
}

func (fs FloorSystem) IsCompleted() bool {
	for _, w := range fs.Works {
		if !w.IsCompleted() {
			return false
		}
	}
	return true
}
