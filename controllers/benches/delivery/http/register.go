package http

import (
	"github.com/gin-gonic/gin"
	"scameiki-and-places/controllers/benches"
)

func RegisterHttpEndPoints(router *gin.RouterGroup, useCase benches.UseCase){
	h:= NewHandler(useCase)

	bechRouter := router.Group("/bench")
	{
		bechRouter.GET("",h.GetAll)
	}
}
