package service

import (
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
)

type ArticleService interface {
	Get(ctx *gin.Context) ([]models.Article, error)
	FindById(ctx *gin.Context, id uint) (models.Article, error)
	Create(ctx *gin.Context, article models.Article) (models.Article, error)
	Update(ctx *gin.Context, id uint, article models.Article) (models.Article, error)
	Delete(ctx *gin.Context, id uint) error
}
