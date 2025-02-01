package handler

import (
	"fmt"
	"gopportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	id := ctx.Query("id")
	ctx.BindJSON(&request)

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id); err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
	}

}
