package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongo/configs"
	"go-gin-mongo/routes"
)

func main() {
	r := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(r)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
