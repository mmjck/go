package add

import (
	"fmt"
	"net/http"
	"task/common"
	"task/data"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func InitHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (handler *Handler) Add(context *gin.Context) {
	var obj data.Task

	err := context.ShouldBindJSON(&obj)

	if err != nil {
		common.BadRequest(context, http.StatusBadRequest, "Request has invalid body",
			[]data.ErrorDetail{
				{
					ErrorType:    data.ErrorTypeError,
					ErrorMessage: "Request has invalid body",
				},
				{
					ErrorType:    data.ErrorTypeValidation,
					ErrorMessage: err.Error(),
				},
			})

		return
	}

	result, errorResponse := handler.service.Add(&obj)
	if errorResponse != nil {
		common.BadRequest(context, http.StatusBadRequest, fmt.Sprintf("Error in Adding task by name %s", obj.Name),
			[]data.ErrorDetail{
				*errorResponse,
			})
	}

	common.Ok(context, http.StatusOK, fmt.Sprintf("successfully Added task with name %s", obj.Name), result)

}
