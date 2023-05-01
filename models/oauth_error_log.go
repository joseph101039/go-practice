package models

import "time"

func init() {
	var _ Model = (*Group)(nil) // 檢查是否實作 Model interface
}

// OauthErrorLog token model
type OauthErrorLog struct {
	BaseModel
	Logged      time.Time `gorm:"->;column:logged;type:timestamp;default:CURRENT_TIMESTAMP();not null"` // read only column
	SystemName  string    `gorm:"column:system_name;not null;size:255;comment:系統名稱"`
	ClientID    string    `gorm:"column:client_id;length:255;not null;comment:用戶ID"`
	RedirectURI string    `gorm:"column:redirect_uri;length:255;not null;comment:重新導向網址"`
	GrantType   string    `gorm:"column:grant_type;type:ENUM('authorization_code','service_token','refresh_token','password');not null;comment:授權模式"` // 參照 oauth2.GrantType 定義
	ErrorCode   string    `gorm:"column:error_code;type:varchar(10);not null;comment:錯誤碼"`
	Data        string    `gorm:"column:data;type:text;not null"`
}

// TableName bind the table name
func (o OauthErrorLog) TableName() string {
	return "oauth_error_log"
}
