package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOppeningHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Funcionou o GO",
	})
}