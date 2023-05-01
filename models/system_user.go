package models

import (
	"time"

	"gorm.io/gorm"
)

// SystemUser 後台系統經過 Google 登入後的紀錄
type SystemUser struct {
	gorm.Model
	Id           uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Name         string         `gorm:"column:name;not null;size:255"`
	Email        string         `gorm:"column:email;index:,length:10;not null;size:255"`
	GoogleUserId uint           `gorm:"column:google_user_id;type:int unsigned;default:null"`
	LastLoginAt  time.Time      `gorm:"column:last_login_at;type:timestamp;not null"`
	CreatedAt    time.Time      `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt    time.Time      `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`
}

func (s SystemUser) TableName() string {
	return "system_users"
}
