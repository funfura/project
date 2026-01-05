package work

import "project/internal/domain/common"

type Work struct {
	ID     common.ID
	Name   string
	Status common.Status
}

func NewWork(id common.ID, name string) Work {
	return Work{
		ID:     id,
		Name:   name,
		Status: common.NotCompleted,
	}
}

func (w *Work) Complete() {
	w.Status = common.Completed
}

func (w Work) IsCompleted() bool {
	return w.Status.IsCompleted()
}
