package services

import (
	"github.com/sithuramyo/go-crud/implementations"
	"github.com/sithuramyo/go-crud/models"
	"gorm.io/gorm"
)

type ScheduleService struct {
	*implementations.CRUD[*models.Schedule]
}

func NewScheduleService(db *gorm.DB) *ScheduleService {
	return &ScheduleService{
		CRUD: implementations.NewCRUD[*models.Schedule](db, []string{"clinic_id"}),
	}
}
