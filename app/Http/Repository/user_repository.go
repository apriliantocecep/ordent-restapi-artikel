package repository

import (
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx *gin.Context, db *gorm.DB, user models.User) (models.User, error)
	FindByEmail(ctx *gin.Context, db *gorm.DB, email string) (models.User, error)
}
