package handler

import (
	"net/http"
	"strings"

	"email.v1/api/dto"
	"email.v1/api/service"
	"email.v1/utils/constanta"
	"email.v1/utils/helper"
	"github.com/gin-gonic/gin"
)



type handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateUser(c *gin.Context) {
	schema := dto.RequestUser{}

	err := c.BindJSON(&schema)
	if err != nil {
		c.JSON(400, helper.ErrorResponse(err.Error()))
		return
	}

	err = h.service.CreateUser(schema)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.JSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.JSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.SuccessResponse("success create user"))
}






