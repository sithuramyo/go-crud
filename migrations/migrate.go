package main

import (
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(
		&models.Clinic{},
		&models.Schedule{},
	)
}
