package models

import (
	"time"

	"gorm.io/gorm"
)

type System struct {
	Id        uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Name      string         `gorm:"column:name;not null;size:255;comment:系統名稱"`
	Alias     string         `gorm:"column:alias;not null;size:255;comment:系統代碼"`
	CreatedAt time.Time      `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt time.Time      `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`
}

// TableName bind the table name
func (system System) TableName() string {
	return "systems"
}
