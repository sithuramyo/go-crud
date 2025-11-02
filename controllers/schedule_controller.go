package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/models"
	"github.com/sithuramyo/go-crud/services"
)

func GetSchedules(c *gin.Context) {
	services.List[models.Schedule](c, initializers.DB, []string{"clinic_id", "date", "Clinic"})
}

func GetSchedule(c *gin.Context) {
	services.Get[models.Schedule](c, initializers.DB)
}

func CreateSchedule(c *gin.Context) {
	services.Create[models.Schedule](c, initializers.DB)
}

func UpdateSchedule(c *gin.Context) {
	services.Update[models.Schedule](c, initializers.DB)
}

func DeleteSchedule(c *gin.Context) {
	services.Delete[models.Schedule](c, initializers.DB)
}
