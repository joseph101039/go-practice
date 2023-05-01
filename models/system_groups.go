package models

import (
	"time"
)

func init() {
	var _ Model = (*SystemGroup)(nil) // 檢查是否實作 Model interface
}

/*
*
當一個 system 成功登入後可以存取的 groups 範圍,
如撈取使用者與組織單位資料
*
*/
type SystemGroup struct {
	BaseModel
	Id        uint      `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	SystemId  uint      `gorm:"column:system_id;type:int unsigned;not null"`
	GroupId   uint      `gorm:"column:group_id;type:int unsigned;not null"`
	CreatedAt time.Time `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;"` // read only column

	// Assocations
	System System `gorm:"ForeignKey:SystemId;AssociationForeignKey:Id"`
	Group  Group  `gorm:"ForeignKey:GroupId;AssociationForeignKey:Id"`
}

// TableName bind the table name
func (s SystemGroup) TableName() string {
	return "system_groups"
}
