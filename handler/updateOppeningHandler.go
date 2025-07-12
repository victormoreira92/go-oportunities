package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/schemas"
)


func UpdateOppeningHandler(ctx *gin.Context){
	request := CreateOppeningUpdate{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil{
		looger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	if request.Company != "" {
		oppening.Company = request.Company
	}

	if request.Role != "" {
		oppening.Role = request.Role
	}

	if request.Link != "" {
		oppening.Link = request.Link
	}

	if request.Location != "" {
		oppening.Location = request.Location
	}

	if request.Remote != nil {
		oppening.Remote = *request.Remote
	}	
	
	if request.Salary > 0 {
		oppening.Salary = request.Salary
	}



	if err := db.Save(&oppening).Error ; err != nil {
		looger.Errorf("Error updating oppening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Errorf("error updating opppening: %v", id).Error())
		return
	}

	sendSuccess(ctx, "updating-oppening", oppening)



}