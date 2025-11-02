package initializers

import (
	"github.com/sithuramyo/go-crud/handlers"
	"github.com/sithuramyo/go-crud/services"
	"gorm.io/gorm"
)

type AppContainer struct {
	ClinicService   *services.ClinicService
	ScheduleService *services.ScheduleService

	ClinicHandler   *handlers.ClinicHandler
	ScheduleHandler *handlers.ScheduleHandler
}

func InitializeServices(db *gorm.DB) *AppContainer {
	clinicService := services.NewClinicService(db)
	scheduleService := services.NewScheduleService(db)

	clinicHandler := handlers.NewClinicHandler(clinicService)
	scheduleHandler := handlers.NewScheduleHandler(scheduleService)

	return &AppContainer{
		ClinicService:   clinicService,
		ScheduleService: scheduleService,
		ClinicHandler:   clinicHandler,
		ScheduleHandler: scheduleHandler,
	}
}
