package router 

import (
	"github.com/gin-gonic/gin"
)

func Initialize (){
	//Inicializando rotas
	router := gin.Default()

	//Inicializa routes
	InitializeRoutes(router)
	
  router.Run(":8080")
}