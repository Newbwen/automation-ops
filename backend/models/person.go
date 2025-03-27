package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Person struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
}

func (p *Person) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}

func (p *Person) TableName() string {
	return "persons"
}
