package services

import (
	"github.com/sithuramyo/go-crud/implementations"
	"github.com/sithuramyo/go-crud/models"
	"gorm.io/gorm"
)

type ClinicService struct {
	*implementations.CRUD[*models.Clinic]
}

// func (s *ClinicService) SpecialLogic() string {
// 	return "Extra logic for clinics"
// }

func NewClinicService(db *gorm.DB) *ClinicService {
	return &ClinicService{
		CRUD: implementations.NewCRUD[*models.Clinic](db, []string{"name", "address"}), // default search fields
	}
}
