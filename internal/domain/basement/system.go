package basement

import (
	"project/internal/domain/common"
	"project/internal/domain/work"
)

type System struct {
	ID    common.ID
	Name  string
	Works []work.Work
}

func (s System) IsCompleted() bool {
	for _, w := range s.Works {
		if !w.IsCompleted() {
			return false
		}
	}
	return true
}
