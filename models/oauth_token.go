package models

import "time"

// OauthToken token model
type OauthToken struct {
	Id            uint   `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	OauthClientId uint   `gorm:"column:oauth_client_id;type:int unsigned;not null;comment:重新導向網址"`
	ClientID      string `gorm:"column:client_id;length:255;not null;comment:用戶ID"`
	RedirectURI   string `gorm:"column:redirect_uri;length:255;not null;comment:重新導向網址"`
	Scope         string `gorm:"column:scope;length:255;default:organization;comment:授權範圍"`

	//CodeChallenge       string        `bson:"CodeChallenge"`
	//CodeChallengeMethod string        `bson:"CodeChallengeMethod"`

	Code             string        `gorm:"column:code;length:255;default:null;index:,length:6;comment:驗證代碼"`
	CodeCreatedAt    *time.Time    `gorm:"column:code_created_at;type:timestamp NULL;comment:驗證碼創立時間"`
	CodeExpiresIn    time.Duration `gorm:"column:code_expires_in;type:bigint;default:null;comment:授權碼存活奈秒"`
	Access           string        `gorm:"column:access;length:255;default:null;index:,length:6;comment:存取碼"`
	AccessCreatedAt  *time.Time    `gorm:"column:access_created_at;type:timestamp NULL;comment:存取碼創立時間"`
	AccessExpiresIn  time.Duration `gorm:"column:access_expires_in;type:bigint;default:null;comment:存取碼存活奈秒;"`
	Refresh          string        `gorm:"column:refresh;length:255;index:,length:6;default:null;comment:更新碼創立時間"`
	RefreshCreatedAt *time.Time    `gorm:"column:refresh_created_at;type:timestamp NULL;更新碼存活奈秒"` // 設定成 pointer 讓 time 預設可以是 null, 而非 zero-time
	RefreshExpiresIn time.Duration `gorm:"column:refresh_expires_in;type:bigint;default:null"`
	OldRefresh       string        `gorm:"column:old_refresh;length:255;default:null;comment:前一次過期的refresh token"`

	GrantType    string `gorm:"column:grant_type;type:ENUM('authorization_code','service_token','refresh_token','password');default:null;comment:授權模式"` // 參照 oauth2.GrantType 定義
	GoogleUserId uint   `gorm:"column:google_user_id;type:int unsigned;default:null"`                                                                   // 如果是 authorization_code 轉跳, 需要紀錄登入的 google 帳號使用者
	IsTesting    bool   `gorm:"column:is_testing;type:bigint;default:0;comment:測試工具用途token,不可打oauth以外API"`

	// Assocations
	OauthClient OauthClient `gorm:"ForeignKey:OauthClientId;AssociationForeignKey:Id"`
	GoogleUser  GoogleUser  `gorm:"ForeignKey:GoogleUserId;AssociationForeignKey:Id"`
}

// TableName bind the table name
func (o OauthToken) TableName() string {
	return "oauth_tokens"
}
