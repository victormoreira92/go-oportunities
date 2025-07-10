package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/handler"
)

func InitializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOppeningHandler)

		v1.GET("/openings", handler.ListOppeningHandler)

		v1.POST("/opening", handler.CreateOppeningHandler)

		v1.DELETE("/opening", handler.DeleteOppeningHandler)

		v1.PUT("/opening", handler.UpdateOppeningHandler)
	}
}