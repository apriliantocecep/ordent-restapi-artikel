package routes

import (
	controller "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Controller"
	middleware "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	authMiddleware middleware.AuthMiddlewareConfig,
	homeController controller.HomeController,
	authController controller.AuthController,
	articleController controller.ArticleController,
) *gin.Engine {
	router := gin.Default()

	router.GET("/", homeController.Index)
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	article := router.Group("/article").Use(authMiddleware.AuthRequired)
	{
		article.GET("/", articleController.Index)
		article.POST("/", articleController.Store)
		article.PUT("/:id", articleController.Update)
		article.DELETE("/:id", articleController.Delete)
		article.GET("/:id", articleController.Show)
	}

	return router
}
