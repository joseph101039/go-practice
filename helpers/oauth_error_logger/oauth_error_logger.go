package oauth_error_logger

import (
	"goroutine/app/models"
	"goroutine/app/repositories"
	"goroutine/helpers/db"
)

type OauthErrorLogger struct {
	record *models.OauthErrorLog
}

var oauthError OauthErrorLogger = OauthErrorLogger{}

func SetClientId(clientID string) OauthErrorLogger {
	system, err := (&repositories.System{}).GetByClientId(clientID)
	if err != nil {
		oauthError.record.SystemName = ""
	} else {
		oauthError.record.SystemName = system.Name
	}

	oauthError.record.ClientID = clientID
	return oauthError
}

func SetRedirectUri(redirectUri string) OauthErrorLogger {
	oauthError.record.RedirectURI = redirectUri
	return oauthError
}

func SetGrantType(grantType string) OauthErrorLogger {
	oauthError.record.GrantType = grantType
	return oauthError
}

func SetErrorCode(errorCode string) OauthErrorLogger {
	oauthError.record.ErrorCode = errorCode
	return oauthError
}

func SetData(data string) OauthErrorLogger {
	oauthError.record.Data = data
	return oauthError
}

// Insert return true if log record is successfully inserted
func Insert() bool {
	tx := db.DB.Create(oauthError.record)
	oauthError = OauthErrorLogger{}
	return tx.Error == nil
}
