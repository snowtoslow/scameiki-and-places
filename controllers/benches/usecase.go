package benches

import (
	"context"
	"scameiki-and-places/models"
)

type UseCase interface {
	GetBenches(ctx context.Context) ([]*models.Bench,error)
	GetBenchById(ctx context.Context,id int)(*models.Bench,error)
}
