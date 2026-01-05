package floorworks

import (
	"project/internal/domain/common"
	"project/internal/domain/work"
)

func NewStandardFloorWorks(id common.ID) FloorSystem {
	return FloorSystem{
		ID:   id,
		Name: "Типовые этажные работы",
		Works: []work.Work{
			work.NewWork("floor-cable-mgn", "Прокладка кабеля МГН"),
			work.NewWork("floor-cable-internet", "Прокладка кабеля интернет"),
			work.NewWork("floor-button-install", "Установка кнопки"),
			work.NewWork("floor-button-connect", "Подключения кнопки"),
			work.NewWork("floor-lump-install", "Установка лампы"),
			work.NewWork("floor-lump-connect", "Подключение лампы"),
			work.NewWork("floor-shield-install", "Установка пожарного щита"),
			work.NewWork("floor-KRT-install", "Установка Крт/Ук"),
			work.NewWork("floor-KRT-connect", "Подключение плинта в крт"),
			work.NewWork("floor-PR-install", "Установка ПР"),
			work.NewWork("floor-PR-connect", "Подключение ПР"),
			work.NewWork("floor-commutator-install", "Установка коммутатора"),
			work.NewWork("floor-commutator-connect", "Подключение коммутатора"),
		},
	}
}
