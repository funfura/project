package uspd

import (
	"project/internal/domain/basement"
	"project/internal/domain/common"
	"project/internal/domain/work"
)

func NewUSPDSystem(id common.ID) basement.System {
	return basement.System{
		ID:   id,
		Name: "УСПД",
		Works: []work.Work{
			work.NewWork("uspd-shield-wiring", "Обвязка щитовой"),
			work.NewWork("uspd-meters", "Подключение счётчиков"),
			work.NewWork("uspd-junction-box", "Разделка кабелей в коробке 100x100"),
			work.NewWork("uspd-install", "Установка УСПД"),
			work.NewWork("uspd-power", "Питание от ЩС"),
			work.NewWork("uspd-internet", "Интернет от УСПД к шкафу Telco"),
			work.NewWork("uspd-connect", "Подключение УСПД"),
			work.NewWork("uspd-parking-unit", "Подключение УСПД к щитовой парковки"),
		},
	}
}
