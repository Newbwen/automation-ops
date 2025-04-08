package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string         `gorm:"type:char(36);not null;primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除
}

// 生成UUID
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}
