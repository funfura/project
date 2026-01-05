package backbone

import (
	"project/internal/domain/basement"
	"project/internal/domain/common"
	"project/internal/domain/work"
)

type CableType string

const (
	UTP10 CableType = "UTP 10x2"
	UTP16 CableType = "UTP 16x2"
)

func CableByApartments(apartments int) CableType {
	if apartments <= 5 {
		return UTP10
	}
	return UTP16
}

func PlintsCount(cable CableType) int {
	if cable == UTP10 {
		return 1
	}
	return 2
}

func NewBackboneSystem(id common.ID, apartments int) basement.System {
	return basement.System{
		ID:   id,
		Name: "Стояковый интернет",
		Works: []work.Work{
			work.NewWork("bb-cabinet", "Установлен техэтажный шкаф"),
			work.NewWork("bb-cable-input", "Кабель заведён в шкаф"),
			work.NewWork("bb-plints-connect", "Плинты расключены"),
		},
	}
}
