package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/controllers"
)

func RegisterAPIRoutes(r *gin.Engine) {
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
	}
}
