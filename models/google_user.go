package models

import (
	"time"

	"gorm.io/gorm"
)

type GoogleUser struct {
	Id          uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Email       string         `gorm:"column:email;index:,length:10;not null;size:255"`
	Account     string         `gorm:"account;size:255"`
	OrgUnitPath string         `gorm:"org_unit_path;not null;size:255"`
	CreatedAt   time.Time      `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt   time.Time      `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`
}

func (googleUser GoogleUser) TableName() string {
	return "google_users"
}
