package photo

import (
	"context"
	"scameiki-and-places/models"
)

type PhotoRepository interface {
	CreatePhoto(ctx context.Context,photo *models.Photo) error
	GetPhotos(ctx context.Context,photo *models.Photo) (*models.Photos,error)
	DeletePhoto(ctx context.Context, id int) error
	UpdatePhoto(ctx context.Context,id int) (*models.Photo,error)
}
