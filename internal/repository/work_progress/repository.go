package repository

import (
	"context"
	"database/sql"

	"project/internal/domain/common"
	"project/internal/domain/work"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// ============================
// Work Types
// ============================

// GetAllWorkTypes возвращает все типы работ как work.Work
func (r *Repository) GetAllWorkTypes(ctx context.Context) ([]work.Work, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT code, name
		FROM work_types
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []work.Work
	for rows.Next() {
		var code string
		var name string
		if err := rows.Scan(&code, &name); err != nil {
			return nil, err
		}
		result = append(result, work.Work{
			ID:     common.ID(code), // приводим string к common.ID
			Name:   name,
			Status: common.NotCompleted, // используем common.Status
		})
	}
	return result, nil
}

// ============================
// Work Progress
// ============================

// GetProgressByObject возвращает все работы и их статус для объекта
func (r *Repository) GetProgressByObject(ctx context.Context, objectID common.ID) ([]work.Work, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT wt.code, wt.name, wp.status
		FROM work_progress wp
		JOIN work_types wt ON wp.work_type_id = wt.id
		WHERE wp.object_id = $1
	`, objectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []work.Work
	for rows.Next() {
		var code string
		var name string
		var status string
		if err := rows.Scan(&code, &name, &status); err != nil {
			return nil, err
		}
		result = append(result, work.Work{
			ID:     common.ID(code),
			Name:   name,
			Status: common.Status(status),
		})
	}
	return result, nil
}

// AddOrUpdateProgress создает запись прогресса или обновляет статус
func (r *Repository) AddOrUpdateProgress(ctx context.Context, objectID common.ID, workID common.ID, objectType string, status common.Status, comment string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO work_progress (object_id, work_type_id, object_type, status, comment)
		SELECT $1, id, $2, $3, $4
		FROM work_types
		WHERE code = $5
		ON CONFLICT (work_type_id, object_id) DO UPDATE
		SET status = EXCLUDED.status,
		    comment = EXCLUDED.comment,
		    updated_at = NOW()
	`, objectID, objectType, string(status), comment, workID)
	return err
}

// MarkAsDone помечает работу как выполненную
func (r *Repository) MarkAsDone(ctx context.Context, objectID common.ID, workID common.ID) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE work_progress wp
		SET status = 'done', updated_at = NOW()
		FROM work_types wt
		WHERE wp.work_type_id = wt.id AND wp.object_id = $1 AND wt.code = $2
	`, objectID, workID)
	return err
}
