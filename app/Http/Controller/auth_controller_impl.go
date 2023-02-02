package controller

import (
	"net/http"

	helper "github.com/apriliantocecep/ordent-restapi-artikel/app/Helper"
	request "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Request"
	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	service "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Service"
	"github.com/gin-gonic/gin"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthControllerImpl(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

// Register godoc
// @Summary register user
// @Schemes
// @Description register a user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=response.RegisterResponse}
// @Failure 400 {object} response.Response
// @Param request body request.RegisterRequest true "register request"
// @Router /register [post]
func (controller *AuthControllerImpl) Register(ctx *gin.Context) {
	var errorOutput []error
	registerRequest := request.RegisterRequest{}

	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		errorOutput = append(errorOutput, err)

		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": helper.ParseError(errorOutput...)[0],
			},
		})
		ctx.Abort()
		return
	}

	res, err := controller.AuthService.Register(ctx, registerRequest)
	if err != nil {
		errorOutput = append(errorOutput, err)

		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": helper.ParseError(errorOutput...)[0],
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}

// Login godoc
// @Summary login user
// @Schemes
// @Description login a user using email and password
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=response.LoginResponse}
// @Failure 400 {object} response.Response
// @Param request body request.LoginRequest true "login request"
// @Router /login [post]
func (controller *AuthControllerImpl) Login(ctx *gin.Context) {
	var errorOutput []error
	loginRequest := request.LoginRequest{}

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		errorOutput = append(errorOutput, err)

		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": helper.ParseError(errorOutput...)[0],
			},
		})
		ctx.Abort()
		return
	}

	res, err := controller.AuthService.Login(ctx, loginRequest)
	if err != nil {
		errorOutput = append(errorOutput, err)

		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": helper.ParseError(errorOutput...)[0],
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
}
