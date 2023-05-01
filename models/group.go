package models

import (
	"time"

	"gorm.io/gorm"
)

func init() {
	var _ Model = (*Group)(nil) // 檢查是否實作 Model interface
}

type Group struct {
	BaseModel
	GroupGoogleUsers []GroupGoogleUser

	Id        uint           `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Name      string         `gorm:"column:name;not null;size:255"`
	GroupKey  string         `gorm:"column:group_key;uniqueIndex;not null;size:255;comment:GCP單位ID"`
	Code      string         `gorm:"column:code;default:null;size:255;after:group_key;comment:代碼,未必同Email"`
	Email     string         `gorm:"column:email;index:,length:10;not null;size:255"`
	ParentId  uint           `gorm:"column:parent_id;index;type:int unsigned;default:null"`
	CreatedAt time.Time      `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt time.Time      `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp NULL"`
}

// TableName bind the table name
func (group Group) TableName() string {
	return "groups"
}
