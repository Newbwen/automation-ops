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

// 用户模型
type Users struct {
	BaseModel
	Username string `gorm:"type:varchar(100);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Email    string `gorm:"type:varchar(255);not null";uniqueIndex json:"email"` // 唯一索引
	//允许为空，否则role为空是这个字段设置为空会报错
	RoleID    string    `gorm:"type:char(36)" json:"role_id"`
	LastLogin time.Time `gorm:"type:datetime(3)";default:null; json:"last_login"`
	//constraint:OnDelete:SET NULL; 表示 role 被删除时，user.role_id 设为 NULL，防止级联删除用户。
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:SET NULL" json:"role"`
}

// 角色模型
type Role struct {
	BaseModel
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

// 生成UUID
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}

// 设置表名
func (u *Users) TableName() string {
	return "user"
}

func (r *Role) TableName() string {
	return "role"
}
