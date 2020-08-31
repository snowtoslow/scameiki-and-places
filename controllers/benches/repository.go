package benches

import (
	"context"
	"scameiki-and-places/models"
)

type BenchRepository interface {
	/*CreateBench(ctx context.Context,bench *models.Bench) error
	UpdateBench(ctx context.Context,id int) (*models.Bench,error)*/
	GetBenches(ctx context.Context) ([]*models.Bench,error)
	GetBenchById(ctx context.Context,id int) (*models.Bench,error)
	DeleteBench(ctx context.Context, id int) error
}
