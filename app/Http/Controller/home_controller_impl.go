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

// Index godoc
// @Summary get home
// @Schemes
// @Description get home welcome response
// @Tags home
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router / [get]
func (controller *HomeControllerImpl) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data: gin.H{
			"message": "Welcome",
		},
	})
}
