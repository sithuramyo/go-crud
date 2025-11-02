package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ClinicID  uuid.UUID `gorm:"type:uuid;not null" json:"clinicId"`
	Clinic    Clinic    `gorm:"foreignKey:ClinicID;references:ID" json:"clinic"`
	MaxSlot   int32     `json:"maxSlot"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Schedule) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
