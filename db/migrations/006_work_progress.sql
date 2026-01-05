BEGIN;

-- ======================
-- Таблица прогресса работ
-- ======================
CREATE TABLE IF NOT EXISTS work_progress (
    object_id UUID NOT NULL,            -- id дома/подъезда/квартиры
    work_code TEXT NOT NULL             -- код работы, как в work_types.code
        REFERENCES work_types(code)
        ON DELETE CASCADE,
    
    object_type TEXT NOT NULL CHECK (
        object_type IN ('basement', 'floor', 'apartment')
    ),

    status TEXT NOT NULL CHECK (
        status IN ('pending', 'done')
    ) DEFAULT 'pending',

    comment TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (object_id, work_code)
);

-- ======================
-- Индексы для скорости
-- ======================
CREATE INDEX IF NOT EXISTS idx_work_progress_object
    ON work_progress (object_type, object_id);

CREATE INDEX IF NOT EXISTS idx_work_progress_status
    ON work_progress (status);

COMMIT;
