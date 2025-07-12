package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victormoreira92/go-oportunities/schemas"
)


func CreateOppeningHandler(ctx *gin.Context){
	request := CreateOppeningRequest{}
	looger.Infof("request received: %+v", request)

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		looger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	oppening := schemas.Oppening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Link: request.Link,
		Salary: request.Salary,
		Remote: *request.Remote,
	}


	if err := db.Create(&oppening).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "error creating opening databse")
		looger.Errorf("error creating: %v", err.Error())
		return
	}
	looger.Info("Oppening created sucessull")
	sendSuccess(ctx, "creating-oppening", oppening)

}