BEGIN;

-- ======================
-- Backbone (Стояковый интернет)
-- ======================
INSERT INTO work_types (system_code, scope, code, name)
VALUES
    ('bb', 'basement', 'bb-cabinet', 'Установлен техэтажный шкаф'),
    ('bb', 'basement', 'bb-cable-input', 'Кабель заведён в шкаф'),
    ('bb', 'basement', 'bb-plints-connect', 'Плинты расключены')
ON CONFLICT (code) DO NOTHING;

-- ======================
-- MGN
-- ======================
INSERT INTO work_types (system_code, scope, code, name)
VALUES
    ('mgn', 'basement', 'mgn-cable-riser', 'Кабель по стояку (UTP 4x2, КПС 2x1.5)'),
    ('mgn', 'basement', 'mgn-cable-between-entrances', 'Кабель между подъездами'),
    ('mgn', 'basement', 'mgn-cable-parking', 'Кабель из парковки'),
    ('mgn', 'basement', 'mgn-cable-cluster', 'Соединение между домами'),
    ('mgn', 'basement', 'mgn-cable-security', 'Соединение с охраной'),
    ('mgn', 'basement', 'mgn-power-input', 'Подключение питания от ЩС'),
    ('mgn', 'basement', 'mgn-psu-install', 'Установка блока питания'),
    ('mgn', 'basement', 'mgn-psu-wiring', 'Расключка блока питания'),
    ('mgn', 'basement', 'mgn-battery', 'Подключение батареи'),
    ('mgn', 'basement', 'mgn-security-unit', 'Установка блока связи в охране'),
    ('mgn', 'basement', 'mgn-check', 'Проверка системы МГН')
ON CONFLICT (code) DO NOTHING;

-- ======================
-- USPD
-- ======================
INSERT INTO work_types (system_code, scope, code, name)
VALUES
    ('uspd', 'basement', 'uspd-shield-wiring', 'Обвязка щитовой'),
    ('uspd', 'basement', 'uspd-meters', 'Подключение счётчиков'),
    ('uspd', 'basement', 'uspd-junction-box', 'Разделка кабелей в коробке 100x100'),
    ('uspd', 'basement', 'uspd-install', 'Установка УСПД'),
    ('uspd', 'basement', 'uspd-power', 'Питание от ЩС'),
    ('uspd', 'basement', 'uspd-internet', 'Интернет от УСПД к шкафу Telco'),
    ('uspd', 'basement', 'uspd-connect', 'Подключение УСПД'),
    ('uspd', 'basement', 'uspd-parking-unit', 'Подключение УСПД к щитовой парковки')
ON CONFLICT (code) DO NOTHING;

-- ======================
-- Floor works
-- ======================
INSERT INTO work_types (system_code, scope, code, name)
VALUES
    ('floor', 'floor', 'floor-cable-mgn', 'Прокладка кабеля МГН'),
    ('floor', 'floor', 'floor-cable-internet', 'Прокладка кабеля интернет'),
    ('floor', 'floor', 'floor-button-install', 'Установка кнопки'),
    ('floor', 'floor', 'floor-button-connect', 'Подключение кнопки'),
    ('floor', 'floor', 'floor-lump-install', 'Установка лампы'),
    ('floor', 'floor', 'floor-lump-connect', 'Подключение лампы'),
    ('floor', 'floor', 'floor-shield-install', 'Установка пожарного щита'),
    ('floor', 'floor', 'floor-KRT-install', 'Установка КРТ/УК'),
    ('floor', 'floor', 'floor-KRT-connect', 'Подключение плинта в КРТ'),
    ('floor', 'floor', 'floor-PR-install', 'Установка ПР'),
    ('floor', 'floor', 'floor-PR-connect', 'Подключение ПР'),
    ('floor', 'floor', 'floor-commutator-install', 'Установка коммутатора'),
    ('floor', 'floor', 'floor-commutator-connect', 'Подключение коммутатора')
ON CONFLICT (code) DO NOTHING;

COMMIT;
