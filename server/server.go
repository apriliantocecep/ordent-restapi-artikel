//go:build wireinject
// +build wireinject

package server

import (
	controller "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Controller"
	middleware "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Middleware"
	repository "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Repository"
	service "github.com/apriliantocecep/ordent-restapi-artikel/app/Http/Service"
	"github.com/apriliantocecep/ordent-restapi-artikel/config"
	"github.com/apriliantocecep/ordent-restapi-artikel/routes"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var homeSet = wire.NewSet(
	controller.NewHomeControllerImpl,
	wire.Bind(new(controller.HomeController), new(*controller.HomeControllerImpl)),
)

var authSet = wire.NewSet(
	repository.NewUserRepositoryImpl,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	service.NewAuthServiceImpl,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	controller.NewAuthControllerImpl,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var articleSet = wire.NewSet(
	repository.NewArticleRepositoryImpl,
	wire.Bind(new(repository.ArticleRepository), new(*repository.ArticleRepositoryImpl)),
	service.NewArticleServiceImpl,
	wire.Bind(new(service.ArticleService), new(*service.ArticleServiceImpl)),
	controller.NewArticleControllerImpl,
	wire.Bind(new(controller.ArticleController), new(*controller.ArticleControllerImpl)),
)

var configSet = wire.NewSet(
	config.NewConfig,
	config.NewDB,
	config.NewJwtWrapper,
)

var middlewareSet = wire.NewSet(
	middleware.NewAuthMiddlewareConfig,
)

func InitializedServer() *gin.Engine {
	wire.Build(
		configSet,
		homeSet,
		authSet,
		articleSet,
		middlewareSet,
		routes.RegisterRoutes,
	)
	return nil
}
