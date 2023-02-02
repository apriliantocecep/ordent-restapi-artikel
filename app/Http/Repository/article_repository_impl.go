package repository

import (
	"errors"

	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleRepositoryImpl struct {
}

func NewArticleRepositoryImpl() *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) Get(ctx *gin.Context, db *gorm.DB) ([]models.Article, error) {
	var articles []models.Article

	if result := db.Find(&articles); result.Error != nil {
		return []models.Article{}, errors.New(result.Error.Error())
	}

	return articles, nil
}

func (repository *ArticleRepositoryImpl) FindById(ctx *gin.Context, db *gorm.DB, id uint) (models.Article, error) {
	var article models.Article

	if result := db.First(&article, id); result.Error != nil {
		return models.Article{}, errors.New("article not found")
	}

	return article, nil
}

func (repository *ArticleRepositoryImpl) Create(ctx *gin.Context, db *gorm.DB, article models.Article) (models.Article, error) {
	var a models.Article

	a.Title = article.Title
	a.Body = article.Body

	if result := db.Create(&a); result.Error != nil {
		return models.Article{}, errors.New(result.Error.Error())
	}

	return a, nil
}

func (repository *ArticleRepositoryImpl) Update(ctx *gin.Context, db *gorm.DB, id uint, article models.Article) (models.Article, error) {
	a, err := repository.FindById(ctx, db, id)
	if err != nil {
		return models.Article{}, errors.New(err.Error())
	}

	a.Title = article.Title
	a.Body = article.Body

	db.Save(&a)

	return a, nil
}

func (repository *ArticleRepositoryImpl) Delete(ctx *gin.Context, db *gorm.DB, id uint) error {
	var article models.Article

	_, err := repository.FindById(ctx, db, id)
	if err != nil {
		return errors.New(err.Error())
	}

	if result := db.Delete(&article, id); result.Error != nil {
		return errors.New(err.Error())
	}

	return nil
}
