package models

import (
	"time"
)

type MailDomainUser struct {
	Id uint `gorm:"column:id;primary_key;type:int unsigned not null auto_increment" json:"id"`

	MailDomainId      uint              `gorm:"column:mail_domain_id;type:int unsigned;not null" json:"mail_domain_id"`
	UnsyncedUserEmail UnsyncedUserEmail `gorm:"ForeignKey:MailDomainId;AssociationForeignKey:Id;constraint:OnDelete:CASCADE" json:"unsynced_user_email"` // MailDomainId 外鍵指向的 models, 設置連集刪除

	Email     string `gorm:"column:email;uniqueIndex;not null;size:255;comment:Email,唯一鍵" json:"email"`

	// API Response 全部內容
	MemberInfo string `gorm:"column:member_info;type:text;default:null;comment:directory.member API json 格式" json:"member_info"`
	UserInfo   string `gorm:"column:user_info;type:text;default:null;comment:directory.user API json 格式" json:"user_info"`

	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null;->" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null;->" json:"updated_at"`
}

// TableName bind the table name
func (list MailDomainUser) TableName() string {
	return "mail_domain_users"
}
