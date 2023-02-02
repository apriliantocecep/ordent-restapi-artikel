package repository

import (
	"errors"

	helper "github.com/apriliantocecep/ordent-restapi-artikel/app/Helper"
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx *gin.Context, db *gorm.DB, user models.User) (models.User, error) {
	var u models.User
	u.Name = user.Name
	u.Email = user.Email
	u.Password = helper.HashPassword(user.Password)

	if result := db.Create(&u); result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	return u, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx *gin.Context, db *gorm.DB, email string) (models.User, error) {
	var user models.User

	if result := db.Where(&models.User{Email: email}).First(&user); result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	return user, nil
}
