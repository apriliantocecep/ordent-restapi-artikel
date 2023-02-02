package controller

import "github.com/gin-gonic/gin"

type HomeController interface {
	Index(ctx *gin.Context)
}
