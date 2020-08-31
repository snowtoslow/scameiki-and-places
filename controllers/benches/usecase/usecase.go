package usecase

import (
	"context"
	"scameiki-and-places/controllers/benches"
	bmpostgres "scameiki-and-places/controllers/benches/repository/postgres"
	"scameiki-and-places/models"
)

type BenchUseCase struct {
	benchRepository benches.BenchRepository
}

func NewBenchUseCase(benchRepository *bmpostgres.BenchRepository) *BenchUseCase {
	return &BenchUseCase{
		benchRepository: benchRepository,
	}
}


func (b BenchUseCase) GetBenches(ctx context.Context) ([]*models.Bench, error) {
	return b.benchRepository.GetBenches(ctx)
}

func (b BenchUseCase) GetBenchById(ctx context.Context,id int) (*models.Bench, error) {
	return b.benchRepository.GetBenchById(ctx,id)
}


