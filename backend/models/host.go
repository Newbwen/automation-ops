package models

type Host struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique;not null"`
	IP     string `gorm:"not null"`
	Status string `gorm:"default:'unknown'"`
}
