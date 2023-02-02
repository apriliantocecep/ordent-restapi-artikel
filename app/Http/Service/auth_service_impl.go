package service

import (
	"errors"

	helper "github.com/apriliantocecep/ordent-restapi-artikel/app/Helper"
	repository "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Repository"
	request "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Request"
	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/apriliantocecep/ordent-restapi-artikel/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Config         *config.Config
	Jwt            helper.JwtWrapper
}

func NewAuthServiceImpl(userRepository repository.UserRepository, db *gorm.DB, c *config.Config, jwt helper.JwtWrapper) *AuthServiceImpl {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Config:         c,
		Jwt:            jwt,
	}
}

func (service *AuthServiceImpl) Register(ctx *gin.Context, req request.RegisterRequest) (response.RegisterResponse, error) {
	var user models.User

	if result := service.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return response.RegisterResponse{}, errors.New("E-Mail already exists")
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password

	u, err := service.UserRepository.Create(ctx, service.DB, user)
	if err != nil {
		return response.RegisterResponse{}, errors.New(err.Error())
	}

	token, err := service.Jwt.GenerateToken(u)
	if err != nil {
		return response.RegisterResponse{}, errors.New(err.Error())
	}

	return response.RegisterResponse{
		User:  u,
		Token: token,
	}, nil
}

func (service *AuthServiceImpl) Login(ctx *gin.Context, req request.LoginRequest) (response.LoginResponse, error) {
	u, err := service.UserRepository.FindByEmail(ctx, service.DB, req.Email)
	if err != nil {
		return response.LoginResponse{}, errors.New(err.Error())
	}

	match := helper.CheckPasswordHash(req.Password, u.Password)
	if !match {
		return response.LoginResponse{}, errors.New("user not found")
	}

	token, err := service.Jwt.GenerateToken(u)
	if err != nil {
		return response.LoginResponse{}, errors.New(err.Error())
	}

	return response.LoginResponse{
		User:  u,
		Token: token,
	}, nil
}

func (service *AuthServiceImpl) Validate(ctx *gin.Context, req request.ValidateRequest) (response.ValidateResponse, error) {
	claims, err := service.Jwt.ValidateToken(req.Token)
	if err != nil {
		return response.ValidateResponse{}, errors.New(err.Error())
	}

	u, err := service.UserRepository.FindByEmail(ctx, service.DB, claims.Email)
	if err != nil {
		return response.ValidateResponse{}, errors.New(err.Error())
	}

	return response.ValidateResponse{
		User: u,
	}, nil
}
