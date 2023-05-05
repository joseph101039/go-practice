package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type BaseModel struct {
	Id uint `json:"id" gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
}

type NormalModel struct {
	BaseModel
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;->"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null;->"`
}

// String converts model to string format
func (m BaseModel) String() string {
	return toString(m)
}

func toString(m any) string {
	s, _ := json.MarshalIndent(m, "", "  ")
	return fmt.Sprintf("%T: %s\n", m, s)
}
