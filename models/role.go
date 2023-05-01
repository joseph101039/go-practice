package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id        uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Name      string         `gorm:"column:name;size:191;comment:角色名稱"`
	Label     string         `gorm:"column:label;size:191;comment:角色標籤"`
	SystemId  uint           `gorm:"column:system_id;type:int unsigned;not null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;->"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null;->"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`

	// Association
	System System `gorm:"ForeignKey:SystemId;AssociationForeignKey:Id"`
}

// TableName bind the table name
func (role Role) TableName() string {
	return "roles"
}
