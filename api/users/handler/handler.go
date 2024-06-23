package handler

import (
	"net/http"
	"strings"

	"email.v1/api/users/dto"
	"email.v1/api/users/service"
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

	err = h.service.Create(schema)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			c.JSON(400, helper.ErrorResponse(err.Error()))
			return
		}

		c.JSON(500, helper.ErrorResponse(err.Error()))
		return
	}

	helper.SuccessReturn(c, http.StatusOK, "success create user")
}
