package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s successfull", op),
		"data": data,
	})
}