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
	db, err := initializers.ConnectToDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(
		&models.Clinic{},
		&models.Schedule{},
	)
}
