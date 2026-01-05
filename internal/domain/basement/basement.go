package basement

import "project/internal/domain/common"

type Basement struct {
	ID      common.ID
	Systems []System
}

func (b Basement) IsCompleted() bool {
	for _, s := range b.Systems {
		if !s.IsCompleted() {
			return false
		}
	}
	return true
}
