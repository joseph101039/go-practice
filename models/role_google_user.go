package models

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	var _ Model = (*RoleGoogleUser)(nil) // 檢查是否實作 Model interface
}

type RoleGoogleUser struct {
	BaseModel
	Id           uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	RoleId       uint           `gorm:"column:role_id;type:int unsigned;not null"`
	GoogleUserId uint           `gorm:"column:google_user_id;type:int unsigned;not null"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;->"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null;->"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`

	// Association
	Role       Role       `gorm:"ForeignKey:RoleId;AssociationForeignKey:Id"`
	GoogleUser GoogleUser `gorm:"ForeignKey:GoogleUserId;AssociationForeignKey:Id"`
}

func (r RoleGoogleUser) TableName() string {
	return "role_google_users"
}
