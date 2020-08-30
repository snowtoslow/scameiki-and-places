package postgres

import (
	"context"
	"database/sql"
	"scameiki-and-places/models"
)

type BenchRepository struct {
	db *sql.DB
}

func NewBenchRepository(db *sql.DB) *BenchRepository{
	return &BenchRepository{
		db: db,
	}
}

/*func (repo BenchRepository) CreateBench(bench *models.Bench) error{
	return repo.db.QueryRow(
		"INSERT INTO benches (geolocation, photo,) VALUES ($1, $2) RETURNING id",
		bench.Geolocation,
		bench.Photo,
	).Scan(&bench.ID)
}


func (repo BenchRepository) FindBenchById(id int)(*models.Bench,error){
		bench:= &models.Bench{}
		if err := repo.db.QueryRow(
			"SELECT id, geolocation, photo FROM benches WHERE id = $1", id).
			Scan(
				&bench.ID,
				&bench.Geolocation,
				&bench.Photo);
			err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		return bench, nil
}*/

func (repo BenchRepository) GetBenches(ctx context.Context)(*models.Benches,error){
	benches:= &models.Benches{}
	if err := repo.db.QueryRowContext(ctx,
		"SELECT * FROM benches ").
		Scan(&benches);
		err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}
	return benches, nil
}