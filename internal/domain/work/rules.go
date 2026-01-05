package work

import "errors"

var ErrAlreadyComleted = errors.New("work already comleted")

func (w *Work) ComleteSafe() error {
	if w.IsCompleted() {
		return ErrAlreadyComleted
	}
	w.Complete()
	return nil
}
