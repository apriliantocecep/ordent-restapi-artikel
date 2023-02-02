package controller

import (
	"net/http"
	"strconv"

	helper "github.com/apriliantocecep/ordent-restapi-artikel/app/Helper"
	request "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Request"
	response "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Response"
	service "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Service"
	models "github.com/apriliantocecep/ordent-restapi-artikel/app/Models"
	"github.com/gin-gonic/gin"
)

type ArticleControllerImpl struct {
	ArticleService service.ArticleService
}

func NewArticleControllerImpl(articleService service.ArticleService) *ArticleControllerImpl {
	return &ArticleControllerImpl{
		ArticleService: articleService,
	}
}

func (controller *ArticleControllerImpl) Index(ctx *gin.Context) {
	articles, err := controller.ArticleService.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   articles,
	})
}

func (controller *ArticleControllerImpl) Store(ctx *gin.Context) {
	var errorOutput []error
	articleRequest := request.ArticleRequest{}

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
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

	article := models.Article{
		Title: articleRequest.Title,
		Body:  articleRequest.Body,
	}
	a, err := controller.ArticleService.Create(ctx, article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   a,
	})
}

func (controller *ArticleControllerImpl) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	var errorOutput []error
	articleRequest := request.ArticleRequest{}

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
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

	article := models.Article{
		Title: articleRequest.Title,
		Body:  articleRequest.Body,
	}
	a, err := controller.ArticleService.Update(ctx, uint(id), article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   a,
	})
}

func (controller *ArticleControllerImpl) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	err := controller.ArticleService.Delete(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

func (controller *ArticleControllerImpl) Show(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.ParseUint(idParam, 10, 32)

	article, err := controller.ArticleService.FindById(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: gin.H{
				"message": err.Error(),
			},
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   article,
	})
}
