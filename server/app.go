package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"scameiki-and-places/controllers/benches"
	bmhttp "scameiki-and-places/controllers/benches/delivery/http"
	bmpostgres "scameiki-and-places/controllers/benches/repository/postgres"
	bcuusecase "scameiki-and-places/controllers/benches/usecase"
	"time"
)

type App struct {
	httpServer *http.Server
	benchUC benches.UseCase

}

func NewApp() *App{
	dbUrl := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.sslmode"),
		)

	db, err := initDB(dbUrl)

	if err!=nil {
		log.Fatalf("Database connection error: %s",err.Error())
	}

	fmt.Println("DATABASE SUCESSFULY CONECTED!",db)

	benchRepo := bmpostgres.NewBenchRepository(db)

	return &App{
		benchUC: bcuusecase.NewBenchUseCase(benchRepo),
	}
}

func (a *App) Run(port string) error{
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api")

	bmhttp.RegisterHttpEndPoints(api, a.benchUC)

	a.httpServer = &http.Server{
		Addr: ":"+port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe();err!=nil {
			log.Fatalf("Failed to listen and serve")
		}
	}()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt,os.Interrupt)

	<-quit

	ctx,shutdown := context.WithTimeout(context.Background(),5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)

}



func initDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil{
		return nil, err
	}

	if err := db.Ping(); err !=nil{
		return nil, err
	}

	return db, nil
}