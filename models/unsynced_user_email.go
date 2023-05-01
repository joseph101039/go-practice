package models

import (
	"time"
)

type UnsyncedUserEmail struct {
	Id         uint      `gorm:"column:id;primary_key;type:int unsigned not null auto_increment" json:"id"`
	MailDomain string    `gorm:"column:mail_domain;size:191;uniqueIndex;comment:使用者 mail的 domain(@後面)" json:"mail_domain"`
	UserCount  uint      `gorm:"column:user_count;type:int unsigned not null" json:"user_count"`
	Memo       string    `gorm:"column:memo;type:text;default:null;comment:備註" json:"memo"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;->" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null;->" json:"updated_at"`
}

// TableName bind the table name
func (list UnsyncedUserEmail) TableName() string {
	return "unsynced_user_emails"
}
