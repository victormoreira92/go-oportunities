package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/schemas"
)


func ShowOppeningHandler(ctx *gin.Context){
	id := ctx.Query("id")

	if id == ""{
		looger.Errorf("error deleting oppening: %v", "id")
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "id is required").Error())
		return
	}

	oppening := schemas.Oppening{}
	if err := db.First(&oppening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("error oppening with %v not found", id))
		return
	} 

	sendSuccess(ctx, "show-opening", oppening)

}