package response

import (
	"log"
	"net/http"
	"rdm/google_organization/helpers/goerror"

	"github.com/gin-gonic/gin"
)

type MessageCallback = func(statusMessage string, messagegPayload interface{}) string

// HttpError 實作 Error() 讓 錯誤資訊可以透過 recover 機制取得
type HttpError struct {
	StatusCode      uint
	MessageCallback MessageCallback // 錯誤訊息的調整, 傳入 status message 與 callback
	MessagePayload  interface{}     // MessageCallback 的傳參
	ReturnData      interface{}
}

// NewHttpError initial a new HttpError
func NewHttpError(StatusCode uint) HttpError {
	return HttpError{
		StatusCode:      StatusCode,
		MessageCallback: nil,
		MessagePayload:  nil,
		ReturnData:      nil,
	}
}

// SetRecoverResponse 設置 gin response context 在 recover 接收到 Error 的錯誤訊息,
// 使用時請在 controller 開頭並加上 defer, Ex: defer SetRecoverResponse(ctx, 101)
// defaultStatusCode 則是當沒有指定 status code 時使用的預設訊息
// recoverCallbacks 註冊 recover 後要執行的回呼
func SetRecoverResponse(context *gin.Context, defaultStatusCode uint, recoverCallbacks ...func()) {
	if err := recover(); err != nil {
		// 1. Recover 開始先執行 回呼函式
		for _, callback := range recoverCallbacks {
			callback()
		}

		// 2. 本地端印出詳細的 trace back
		log.Println(goerror.GetStackTrace(err)) // print trace back

		var res Response
		// 3. 判斷 error 種類
		switch e := err.(type) {

		case HttpError:
			res = ResponseMaker(e.StatusCode, nil)
			res.setStatusMessage(e.Error())
		case error:
			res = ResponseMaker(defaultStatusCode, e.Error())
		case string:
			res = ResponseMaker(defaultStatusCode, e)
		default:
			res = ResponseMaker(defaultStatusCode, nil)
		}

		context.JSON(http.StatusBadRequest, res)
		return
	}
}

// AppendMessage 在原本的 status message 後面串上新的訊息
func (err *HttpError) AppendMessage(appendMessage string) {
	if appendMessage != "" {
		err.SetMessageCallback(appendMessage, func(statusMessage string, appendMessage interface{}) string {
			return statusMessage + appendMessage.(string)
		})
	}
}

// SetMessageCallback 設置 訊息 callback
// Example:	err := NewHttpError(101).SetMessageCallback(...); panic(err);
func (err *HttpError) SetMessageCallback(messagePayload interface{}, messageCallback MessageCallback) {
	err.MessagePayload = messagePayload
	err.MessageCallback = messageCallback
}

func (err *HttpError) SetReturnData(returnData interface{}) *HttpError {
	err.ReturnData = returnData
	return err
}

// Error implement error interface which returns the error message
func (err HttpError) Error() string {
	// 取得原始錯誤訊息
	message := getResponseMessage(uint(err.StatusCode))
	// 經過 callback 調整訊息內容
	if err.MessageCallback != nil {
		return err.MessageCallback(message, err.MessagePayload)
	} else {
		return message
	}
}
