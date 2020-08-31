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

func (repo BenchRepository) DeleteBench(ctx context.Context,id int)error{

	if _,err := repo.db.Exec("DELETE FROM benches WHERE id=$1;", id);err!=nil{
		return err
	}

	return nil
}

func (repo BenchRepository) CreateBench(ctx context.Context,bench *models.Bench)(*models.Bench,error){
	err := repo.db.QueryRow("INSERT INTO benches(geolocation,photo) VALUES ($1, $2) RETURNING id",bench.Geolocation,bench.Photo).Scan(&bench.ID)
	if err!=nil {
		return nil, err
	}
	return bench,nil
}
