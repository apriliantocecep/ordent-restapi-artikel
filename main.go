package main

import (
	"github.com/apriliantocecep/ordent-restapi-artikel/server"
)

func main() {
	handler := server.InitializedServer()
	handler.Run()
}
