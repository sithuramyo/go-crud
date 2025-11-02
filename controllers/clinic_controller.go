package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/models"
	"github.com/sithuramyo/go-crud/services"
)

func GetClinics(c *gin.Context) {
	services.List[models.Clinic](c, initializers.DB, []string{"name", "address"})
}

func GetClinic(c *gin.Context) {
	services.Get[models.Clinic](c, initializers.DB)
}

func CreateClinic(c *gin.Context) {
	services.Create[models.Clinic](c, initializers.DB)
}

func UpdateClinic(c *gin.Context) {
	services.Update[models.Clinic](c, initializers.DB)
}

func DeleteClinic(c *gin.Context) {
	services.Delete[models.Clinic](c, initializers.DB)
}
