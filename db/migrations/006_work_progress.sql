-- 006_work_progress.sql
BEGIN;

-- ======================
-- Таблица прогресса работ
-- ======================
CREATE TABLE IF NOT EXISTS work_progress (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    work_type_id UUID NOT NULL
        REFERENCES work_types(id)
        ON DELETE CASCADE,

    object_id UUID NOT NULL,
    object_type TEXT NOT NULL CHECK (
        object_type IN ('basement', 'floor', 'apartment')
    ),

    status TEXT NOT NULL CHECK (
        status IN ('pending', 'done')
    ) DEFAULT 'pending',

    comment TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (work_type_id, object_id)
);

-- ======================
-- Индексы для скорости
-- ======================
CREATE INDEX IF NOT EXISTS idx_work_progress_object
    ON work_progress (object_type, object_id);

CREATE INDEX IF NOT EXISTS idx_work_progress_status
    ON work_progress (status);

COMMIT;
