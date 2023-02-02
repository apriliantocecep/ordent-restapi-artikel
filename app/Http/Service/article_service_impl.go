package service

import (
	"errors"

	repository "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Repository"
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	DB                *gorm.DB
}

func NewArticleServiceImpl(articleRepository repository.ArticleRepository, db *gorm.DB) *ArticleServiceImpl {
	return &ArticleServiceImpl{
		ArticleRepository: articleRepository,
		DB:                db,
	}
}

func (service *ArticleServiceImpl) Get(ctx *gin.Context) ([]models.Article, error) {
	articles, err := service.ArticleRepository.Get(ctx, service.DB)
	if err != nil {
		return []models.Article{}, errors.New(err.Error())
	}

	return articles, nil
}

func (service *ArticleServiceImpl) FindById(ctx *gin.Context, id uint) (models.Article, error) {
	article, err := service.ArticleRepository.FindById(ctx, service.DB, id)
	if err != nil {
		return models.Article{}, errors.New(err.Error())
	}

	return article, nil
}

func (service *ArticleServiceImpl) Create(ctx *gin.Context, article models.Article) (models.Article, error) {
	a, err := service.ArticleRepository.Create(ctx, service.DB, article)
	if err != nil {
		return models.Article{}, errors.New(err.Error())
	}

	return a, nil
}

func (service *ArticleServiceImpl) Update(ctx *gin.Context, id uint, article models.Article) (models.Article, error) {
	article, err := service.ArticleRepository.Update(ctx, service.DB, id, article)
	if err != nil {
		return models.Article{}, errors.New(err.Error())
	}

	return article, nil
}

func (service *ArticleServiceImpl) Delete(ctx *gin.Context, id uint) error {
	err := service.ArticleRepository.Delete(ctx, service.DB, id)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
