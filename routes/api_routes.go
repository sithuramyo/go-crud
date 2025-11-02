package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/handlers"
)

type AppHandlers struct {
	ClinicHandler   *handlers.ClinicHandler
	ScheduleHandler *handlers.ScheduleHandler
}

func RegisterAPIRoutes(r *gin.Engine, h *AppHandlers) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Backend API is running"})
	})

	api := r.Group("/api")
	{
		clinic := api.Group("/clinics")
		{
			clinic.GET("/", h.ClinicHandler.List)
			clinic.POST("/", h.ClinicHandler.Create)
			clinic.GET("/:id", h.ClinicHandler.Get)
			clinic.PUT("/:id", h.ClinicHandler.Update)
			clinic.DELETE("/:id", h.ClinicHandler.Delete)
		}

		schedule := api.Group("/schedules")
		{
			schedule.GET("/", h.ScheduleHandler.List)
			schedule.POST("/", h.ScheduleHandler.Create)
			schedule.GET("/:id", h.ScheduleHandler.Get)
			schedule.PUT("/:id", h.ScheduleHandler.Update)
			schedule.DELETE("/:id", h.ScheduleHandler.Delete)
		}
	}
}
