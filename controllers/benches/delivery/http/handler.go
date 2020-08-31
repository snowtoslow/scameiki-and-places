package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scameiki-and-places/controllers/benches"
	"scameiki-and-places/models"
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
	ctx.JSON(http.StatusOK,myBenches)
}

func (handler *Handler) GetBenchById(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":benches.ErrInvalidId})

	}
	myBench, err := handler.useCase.GetBenchById(ctx.Request.Context(),id)
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

	ctx.JSON(http.StatusOK,myBench)
}

func (handler *Handler) DeleteBench(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}

/*	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,err)
		return
	}
*/
	ctx.JSON(http.StatusNoContent,handler.useCase.DeleteBench(ctx.Request.Context(),id))
}

func (handler *Handler) CreateBench(ctx *gin.Context){

	myBench:= models.Bench{}

	if err := ctx.Bind(&myBench);err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":"failed","error":err.Error()})
		return
	}

	createdBench,err:= handler.useCase.CreateBench(ctx.Request.Context(),&myBench)
	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":"failed","error":err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated,&createdBench)

}

func (handler *Handler) UpdateBench(ctx *gin.Context){
	id, err := strconv.Atoi(ctx.Param("id"))

	if err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"failed","error":benches.ErrInvalidId})
		return
	}

	myBench:= models.Bench{
		ID: id,
	}


	if err := ctx.Bind(&myBench);err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{"status":"failed","error":err.Error()})
		return
	}

	updatedBench,err := handler.useCase.UpdateBench(ctx.Request.Context(),&myBench)

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"failed","error":err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent,&updatedBench)
}
