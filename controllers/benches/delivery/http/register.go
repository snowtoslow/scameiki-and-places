package http

import (
	"github.com/gin-gonic/gin"
	"scameiki-and-places/controllers/benches"
)

func RegisterHttpEndPoints(router *gin.RouterGroup, useCase benches.UseCase){
	h:= NewHandler(useCase)

	benchRouter := router.Group("/bench")
	{
		benchRouter.GET("",h.GetAll)
		benchRouter.GET("/:id",h.GetBenchById)
	}
}
