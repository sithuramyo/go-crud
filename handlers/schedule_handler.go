package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sithuramyo/go-crud/services"
)

type ScheduleHandler struct {
	Service *services.ScheduleService
}

func NewScheduleHandler(service *services.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{Service: service}
}

func (h *ScheduleHandler) List(c *gin.Context) {
	h.Service.List(c)
}

func (h *ScheduleHandler) Get(c *gin.Context) {
	h.Service.Get(c)
}

func (h *ScheduleHandler) Create(c *gin.Context) {
	h.Service.Create(c)
}

func (h *ScheduleHandler) Update(c *gin.Context) {
	h.Service.Update(c)
}

func (h *ScheduleHandler) Delete(c *gin.Context) {
	h.Service.Delete(c)
}
