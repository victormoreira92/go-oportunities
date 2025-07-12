package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/schemas"
)

func ListOppeningHandler(ctx *gin.Context){
	list_oppenings := []schemas.Oppening{}

	if err := db.Find(&list_oppenings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listenig oppenings")
		return
	}

	sendSuccess(ctx, "list oppenings", list_oppenings)
}