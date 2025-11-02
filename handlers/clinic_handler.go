package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/services"
)

type ClinicHandler struct {
	Service *services.ClinicService
}

func NewClinicHandler(service *services.ClinicService) *ClinicHandler {
	return &ClinicHandler{Service: service}
}

func (h *ClinicHandler) List(c *gin.Context) {
	h.Service.List(c)
}

func (h *ClinicHandler) Get(c *gin.Context) {
	h.Service.Get(c)
}

func (h *ClinicHandler) Create(c *gin.Context) {
	h.Service.Create(c)
}

func (h *ClinicHandler) Update(c *gin.Context) {
	h.Service.Update(c)
}

func (h *ClinicHandler) Delete(c *gin.Context) {
	h.Service.Delete(c)
}
