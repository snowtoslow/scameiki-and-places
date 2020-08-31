package postgres

import (
	"context"
	"database/sql"
	"log"
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


func (repo BenchRepository) GetBenchById(ctx context.Context,id int)(*models.Bench,error){
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
}

func (repo BenchRepository) GetBenches(ctx context.Context)([]*models.Bench,error){

	rows, err := repo.db.Query("SELECT * FROM benches")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var bks []*models.Bench
	for rows.Next() {
		bk := new(models.Bench)
		err := rows.Scan(&bk.ID, &bk.Photo, &bk.Geolocation)
		if err != nil {
			log.Fatal(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	 benches := bks

	return benches, err

}