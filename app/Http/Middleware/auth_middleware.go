package middleware

import (
	"net/http"
	"strings"

	request "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Request"
	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	service "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Service"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	AuthService service.AuthService
}

func NewAuthMiddlewareConfig(authService service.AuthService) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{
		AuthService: authService,
	}
}

func (middleware *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		})
		ctx.Abort()
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data: gin.H{
				"message": "invalid token",
			},
		})
		ctx.Abort()
		return
	}

	res, err := middleware.AuthService.Validate(ctx, request.ValidateRequest{
		Token: token[1],
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.Set("userId", res.User.ID)
	ctx.Next()
}
