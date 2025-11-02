package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/controllers"
)

func RegisterAPIRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Backend API is running"})
	})
	api := r.Group("/api")
	{
		clinic := api.Group("/clinics")
		{
			clinic.GET("/", controllers.GetClinics)
			clinic.POST("/", controllers.CreateClinic)
			clinic.GET("/:id", controllers.GetClinic)
			clinic.PUT("/:id", controllers.UpdateClinic)
			clinic.DELETE("/:id", controllers.DeleteClinic)
		}
		schedule := api.Group("/schedules")
		{
			schedule.GET("/", controllers.GetSchedules)
			schedule.POST("/", controllers.CreateSchedule)
			schedule.GET("/:id", controllers.GetSchedule)
			schedule.PUT("/:id", controllers.UpdateSchedule)
			schedule.DELETE("/:id", controllers.DeleteSchedule)
		}
	}
}
