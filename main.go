package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/middlewares"
	"github.com/sithuramyo/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	db, err := initializers.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.Use(middlewares.GlobalErrorHandler())
	app := initializers.InitializeServices(db)
	routes.RegisterAPIRoutes(r, &routes.AppHandlers{
		ClinicHandler:   app.ClinicHandler,
		ScheduleHandler: app.ScheduleHandler,
	})

	r.Run()
}
