package usecase

import (
	"context"
	"scameiki-and-places/controllers/benches"
	"scameiki-and-places/models"
)

type BenchUseCase struct {
	benchRepository benches.BenchRepository
}

func NewBenchUseCase(benchRepository benches.BenchRepository) *BenchUseCase{
	return &BenchUseCase{
		benchRepository: benchRepository,
	}
}


func (b BenchUseCase) GetBenches(ctx context.Context) (*models.Benches, error) {
	return b.benchRepository.GetBenches(ctx)
}


