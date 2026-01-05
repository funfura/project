package mgn

import (
	"project/internal/domain/basement"
	"project/internal/domain/common"
	"project/internal/domain/work"
)

func NewMGNSystem(id common.ID) basement.System {
	return basement.System{
		ID:   id,
		Name: "МГН",
		Works: []work.Work{
			work.NewWork("mgn-cable-riser", "Кабель по стояку (UTP 4x2, КПС 2x1.5)"),
			work.NewWork("mgn-cable-between-entrances", "Кабель между подъездами"),
			work.NewWork("mgn-cable-parking", "Кабель из парковки"),
			work.NewWork("mgn-cable-cluster", "Соединение между домами"),
			work.NewWork("mgn-cable-security", "Соединение с охраной"),
			work.NewWork("mgn-power-input", "Подключение питания от ЩС"),
			work.NewWork("mgn-psu-install", "Установка блока питания"),
			work.NewWork("mgn-psu-wiring", "Расключка блока питания"),
			work.NewWork("mgn-battery", "Подключение батареи"),
			work.NewWork("mgn-security-unit", "Установка блока связи в охране"),
			work.NewWork("mgn-check", "Проверка системы МГН"),
		},
	}
}
