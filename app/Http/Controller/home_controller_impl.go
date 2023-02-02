package controller

import (
	"net/http"

	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	"github.com/gin-gonic/gin"
)

type HomeControllerImpl struct {
}

func NewHomeControllerImpl() *HomeControllerImpl {
	return &HomeControllerImpl{}
}

func (controller *HomeControllerImpl) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data: gin.H{
			"message": "Welcome",
		},
	})
}
