package routes

import (
	controller "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Controller"
	middleware "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Middleware"
	docs "github.com/apriliantocecep/ordent-restapi-artikel/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(
	authMiddleware middleware.AuthMiddlewareConfig,
	homeController controller.HomeController,
	authController controller.AuthController,
	articleController controller.ArticleController,
) *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	return router
}
