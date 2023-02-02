package repository

import (
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Get(ctx *gin.Context, db *gorm.DB) ([]models.Article, error)
	FindById(ctx *gin.Context, db *gorm.DB, id uint) (models.Article, error)
	Create(ctx *gin.Context, db *gorm.DB, article models.Article) (models.Article, error)
	Update(ctx *gin.Context, db *gorm.DB, id uint, article models.Article) (models.Article, error)
	Delete(ctx *gin.Context, db *gorm.DB, id uint) error
}
