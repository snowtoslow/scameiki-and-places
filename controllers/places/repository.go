package places

import (
	"context"
	"scameiki-and-places/models"
)

type PlaceRepository interface {
	CreatePlace(ctx context.Context,place *models.Place) error
	GetPlaces(ctx context.Context) (*models.Places,error)
	DeletePlace(ctx context.Context, id int) error
	GetPlaceById(ctx context.Context,id int) (*models.Place,error)
	UpdatePlace(ctx context.Context,id int) (*models.Place,error)
}
