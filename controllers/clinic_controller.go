package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/initializers"
	"github.com/sithuramyo/go-crud/models"
	"github.com/sithuramyo/go-crud/services"
)

func GetClinics(c *gin.Context) {
	services.GenericList[models.Clinic](c, initializers.DB, []string{"name", "address"})
}

func GetClinic(c *gin.Context) {
	services.GenericGet[models.Clinic](c, initializers.DB)
}

func CreateClinic(c *gin.Context) {
	services.GenericCreate[models.Clinic](c, initializers.DB)
}

func UpdateClinic(c *gin.Context) {
	services.GenericUpdate[models.Clinic](c, initializers.DB)
}

func DeleteClinic(c *gin.Context) {
	services.GenericDelete[models.Clinic](c, initializers.DB)
}
