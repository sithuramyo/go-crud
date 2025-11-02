package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.RegisterAPIRoutes(r)
	r.Run()
}
