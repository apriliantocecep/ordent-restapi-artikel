package service

import (
	request "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Request"
	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register(ctx *gin.Context, req request.RegisterRequest) (response.RegisterResponse, error)
	Login(ctx *gin.Context, req request.LoginRequest) (response.LoginResponse, error)
	Validate(ctx *gin.Context, req request.ValidateRequest) (response.ValidateResponse, error)
}
