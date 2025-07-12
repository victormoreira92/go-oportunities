package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/handler"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	docs "github.com/victormoreira92/go-oportunities/docs"
)

func InitializeRoutes(router *gin.Engine){
	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath

	handler.InitializeHandler()
	
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOppeningHandler)

		v1.GET("/openings", handler.ListOppeningHandler)

		v1.POST("/opening", handler.CreateOppeningHandler)

		v1.DELETE("/opening", handler.DeleteOppeningHandler)

		v1.PUT("/opening", handler.UpdateOppeningHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}