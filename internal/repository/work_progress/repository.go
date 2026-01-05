package workprogress

import (
	"context"
	"project/internal/domain/work"
)

type Repository interface {
	InitForObject(
		ctx context.Context,
		objectType string,
		objectID work.objectID,
		workTypeIDs []work.TypeID,
	)
}
