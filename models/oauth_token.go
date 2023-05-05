package models

import (
	"time"
)

func init() {
	var _ Model = (*OauthToken)(nil) // 檢查是否實作 Model interface
}

// OauthToken token model
type OauthToken struct {
	BaseModel

	// Id            uint   `json:"id,omitempty" gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	OauthClientId uint   `json:"oauth_client_id,omitempty" gorm:"column:oauth_client_id;type:int unsigned;not null;comment:重新導向網址"`
	ClientID      string `json:"client_id,omitempty" gorm:"column:client_id;length:255;not null;comment:用戶ID"`
	RedirectURI   string `json:"redirect_uri,omitempty" gorm:"column:redirect_uri;length:255;not null;comment:重新導向網址"`
	Scope         string `json:"scope" gorm:"column:scope;length:255;default:organization;comment:授權範圍"`

	Code             *string        `json:"code,omitempty" gorm:"column:code;length:255;default:null;index:,length:6;comment:驗證代碼"`
	CodeCreatedAt    *time.Time     `json:"code_created_at,omitempty" gorm:"column:code_created_at;type:timestamp NULL;comment:驗證碼創立時間"`
	CodeExpiresIn    *time.Duration `json:"code_expires_in,omitempty" gorm:"column:code_expires_in;type:bigint;default:null;comment:授權碼存活奈秒"`
	Access           *string        `json:"access,omitempty" gorm:"column:access;length:255;default:null;index:,length:6;comment:存取碼"`
	AccessCreatedAt  *time.Time     `json:"access_created_at,omitempty" gorm:"column:access_created_at;type:timestamp NULL;comment:存取碼創立時間"`
	AccessExpiresIn  *time.Duration `json:"access_expire_in,omitempty" gorm:"column:access_expires_in;type:bigint;default:null;comment:存取碼存活奈秒;"`
	Refresh          *string        `json:"refresh,omitempty" gorm:"column:refresh;length:255;index:,length:6;default:null;comment:更新碼創立時間"`
	RefreshCreatedAt *time.Time     `json:"refresh_created_at,omitempty" gorm:"column:refresh_created_at;type:timestamp NULL;更新碼存活奈秒"` // 設定成 pointer 讓 time 預設可以是 null, 而非 zero-time
	RefreshExpiresIn *time.Duration `json:"refresh_expires_in,omitempty" gorm:"column:refresh_expires_in;type:bigint;default:null"`
	OldRefresh       *string        `json:"old_refresh,omitempty" gorm:"column:old_refresh;length:255;default:null;comment:前一次過期的refresh token"`

	GrantType    string `json:"grant_type,omitempty" gorm:"column:grant_type;type:ENUM('authorization_code','service_token','refresh_token','password');default:null;comment:授權模式"` // 參照 oauth2.GrantType 定義
	GoogleUserId uint   `json:"google_user_id,omitempty" gorm:"column:google_user_id;type:int unsigned;default:null"`                                                               // 如果是 authorization_code 轉跳, 需要紀錄登入的 google 帳號使用者
	IsTesting    bool   `json:",omitempty" gorm:"column:is_testing;type:bigint;default:0;comment:測試工具用途token,不可打oauth以外API"`

	// Assocations
	OauthClient *OauthClient `json:"oauth_client,omitempty" gorm:"ForeignKey:OauthClientId;AssociationForeignKey:Id"`
	GoogleUser  *GoogleUser  `json:"google_user,omitempty" gorm:"ForeignKey:GoogleUserId;AssociationForeignKey:Id"`
}

// TableName bind the table name
func (o OauthToken) TableName() string {

	return "oauth_tokens"
}

// String converts model to string format
func (m OauthToken) String() string {
	return toString(m)
}
