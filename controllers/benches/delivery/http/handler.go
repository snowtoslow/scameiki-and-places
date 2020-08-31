package http

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"scameiki-and-places/controllers/benches"
	"strconv"
)

type Bench struct {
	ID int `json:"id"`
	Geolocation string `json:"geolocation"`
	Photo string `json:"photo"`
}

type Handler struct {
	useCase benches.UseCase
}

func NewHandler(useCase benches.UseCase) *Handler{
	return &Handler{
		useCase: useCase,
	}
}



func (handler *Handler) GetAll(ctx *gin.Context){

	myBenches, err := handler.useCase.GetBenches(ctx.Request.Context())
	if err!=nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer log.Println(handler.useCase.GetBenches(ctx.Request.Context()))
	ctx.JSON(http.StatusOK,myBenches)
}

func (handler *Handler) GetBenchById(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	}
	myBench, err := handler.useCase.GetBenchById(ctx.Request.Context(),id)
	if err!=nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK,myBench)
}
