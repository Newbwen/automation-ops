package models

// 主机模型
type Host struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	IP          string `gorm:"type:varchar(100);not null" json:"ip"`
	Port        int    `gorm:"type:int;not null" json:"port"`
	Status      string `gorm:"type:varchar(100);not null" json:"status"`
	Os          string `gorm:"type:varchar(100);not null" json:"os"`
	Username    string `gorm:"type:varchar(100);not null" json:"username"`
	Password    string `gorm:"type:varchar(100);not null" json:"password"`
	Description string `gorm:"type:varchar(255);default:''" json:"description"`
}
