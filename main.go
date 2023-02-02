package main

import (
	"github.com/apriliantocecep/ordent-restapi-artikel/server"
)

// @title           Ordent Article Online API
// @version         1.0
// @description     Article Online API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Cecep Aprilianto
// @contact.url    https://apriliantocecep.github.io
// @contact.email  cecepaprilianto@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	handler := server.InitializedServer()
	handler.Run()
}
