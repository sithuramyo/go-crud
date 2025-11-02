package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Clinic struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string     `json:"name"`
	Address       string     `json:"address"`
	ContactNumber string     `json:"contactNumber"`
	Schedules     []Schedule `gorm:"foreignKey:ClinicID;references:ID" json:"schedules"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
}

func (c *Clinic) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
