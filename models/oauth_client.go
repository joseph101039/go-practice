package models

import (
	"encoding/json"
	"goroutine/helpers/goerror"
	"log"
	"net/url"
	"time"
)

type OauthClient struct {
	Id           uint      `gorm:"column:id;primary_key;type:int unsigned not null auto_increment"`
	Domain       string    `gorm:"column:domain;length:255;default:null;comment:重新導向網址的網域 json string 格式"`
	SystemId     uint      `gorm:"column:system_id;type:int unsigned;comment:所屬系統"`
	ClientID     string    `gorm:"column:client_id;length:255;not null;comment:用戶ID"`
	ClientSecret string    `gorm:"column:client_secret;size:255;comment:oauth密鑰"`
	Token        string    `gorm:"column:token;size:255;default:null;comment:token"`
	PrivateKey   string    `gorm:"column:private_key;type:text;default:null;comment:RSA私鑰"`
	PublicKey    string    `gorm:"column:public_key;type:text;default:null;comment:RSA公鑰"`
	CreatedAt    time.Time `gorm:"->;column:created_at;type:timestamp;default:CURRENT_TIMESTAMP();not null"`                             // read only column
	UpdatedAt    time.Time `gorm:"->;column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP;not null"` // read only column

	// Assocations
	System System `gorm:"ForeignKey:SystemId;AssociationForeignKey:Id"`
}

// TableName bind the table name
func (o OauthClient) TableName() string {
	return "oauth_clients"
}

// isUriSameHostWithDomain 傳入的 uri 是否與 domain 清單中任一有相同 host
func (o OauthClient) IsUriSameHostWithDomain(uri string) bool {
	var domainList []string
	err := json.Unmarshal([]byte(o.Domain), &domainList)
	goerror.Fatal(err)

	var uriHost string
	if uriUrl, uriErr := url.Parse(uri); uriErr != nil {
		return false
	} else {
		uriHost = uriUrl.Hostname()
	}

	for _, domain := range domainList {
		// 如果有任一 domain match uri 的 hostname 就算相符合
		if domainUrl, domainErr := url.Parse(domain); domainErr == nil {
			if domainUrl.Hostname() == uriHost {
				return true
			}
		}
	}

	// domain 不相符 印出除錯
	log.Printf("uriHost = '%s', domainList = %#v", uriHost, domainList)

	return false
}
