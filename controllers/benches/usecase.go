package benches

import (
	"context"
	"scameiki-and-places/models"
)

type UseCase interface {
	GetBenches(ctx context.Context) ([]*models.Bench,error)
	GetBenchById(ctx context.Context,id int)(*models.Bench,error)
	DeleteBench(ctx context.Context, id int) error
	CreateBench(ctx context.Context,bench *models.Bench)(*models.Bench,error)
}
