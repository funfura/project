package common

type Status int

const (
	NotCompleted Status = iota
	Completed
)

func (s Status) IsCompleted() bool {
	return s == Completed
}
