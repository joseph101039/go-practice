package models

import (
	"time"
)

type GroupGoogleUser struct {
	Id           uint       `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	GroupId      uint       `gorm:"column:group_id;type:int unsigned;not null"`
	Group        Group      `gorm:"ForeignKey:GroupId;AssociationForeignKey:Id"`
	GoogleUserId uint       `gorm:"column:google_user_id;type:int unsigned;not null"`
	GoogleUser   GoogleUser `gorm:"ForeignKey:GoogleUserId;AssociationForeignKey:Id"`
	CreatedAt    time.Time  `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt    time.Time  `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column
}

func (groupUser GroupGoogleUser) TableName() string {
	return "group_google_users"
}
