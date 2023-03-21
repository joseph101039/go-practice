package response

import (
	"fmt"
	"rdm/google_organization/helpers/env"
)

const DEPARTMENT_CODE = "22"
const PROJECT_CODE = "013"

// type Response = map[string]interface{}

type Response struct {
	StatusCode    string      `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	ReturnData    interface{} `json:"return_data"`
}

func (res Response) GetStatusCode() string {
	return res.StatusCode
}

func (res *Response) SetReturnData(data interface{}) *Response {
	res.ReturnData = data
	return res
}
func (res *Response) setStatusMessage(message string) *Response {
	res.StatusMessage = message
	return res
}

// ErrRecoverResponse - deprecated
func ErrRecoverResponse(statusCode uint, returnData interface{}) Response {

	if err, ok := returnData.(error); ok {
		return ResponseMaker(statusCode, err.Error())
	}
	return ResponseMaker(statusCode, returnData)
}

// ResponseMaker 創建 API json response 格式
func ResponseMaker(statusCode uint, returnData interface{}) Response {
	response := Response{
		StatusCode:    fmt.Sprintf("%s%s%d", DEPARTMENT_CODE, PROJECT_CODE, statusCode),
		StatusMessage: getResponseMessage(statusCode),
		ReturnData:    returnData,
	}

	return response
}

func getResponseMessage(statusCode uint) string {

	var message string = ""

	switch statusCode {
	// 200 Status OK: 1, 2, 3, 4 開頭分別為 CRUD
	case 102:
		message = "新增成功"

	case 105:
		message = "取得授權碼成功"

	case 106:
		message = "取得存取碼成功"

	case 201:
		message = "查詢成功"

	case 202:
		message = "無資料"

	case 301:
		message = "更新成功"

	case 401:
		message = "刪除成功"

	// 無法處理的錯誤
	case 502: // 404 not found
		message = "API 不存在"
	case 505:
		message = "存取拒絕: invalid access token"
	case 506:
		message = "存取拒絕: header 未傳入 access-token"
	case 507:
		message = "存取拒絕: access token 已過期"
	case 508:
		message = "存取拒絕: header 未傳入 token"

	case 509:
		message = "後臺登入已經過期, 請重新登入"

	// 400 Bad request: 6, 7, 8, 9 開頭分別為 CRUD
	// Oauth - Generate Authorization Code: see https://datatracker.ietf.org/doc/html/rfc6749#section-4.1.2.1
	case 603:
		message = "server_error:取得授權碼失敗, 請聯絡客服"

	case 605:
		message = "invalid_request"

	case 606:
		message = "unauthorized_client"

	case 607:
		message = "access_denied"

	case 608:
		message = "unsupported_response_type"

	case 609:
		message = "invalid_scope"

	// Oauth - Generate Token: see https://datatracker.ietf.org/doc/html/rfc6749#section-5.2
	case 611:
		message = "server_error:取得授權碼失敗, 請聯絡客服"

	case 612:
		message = "invalid_request"

	case 613:
		message = "invalid_client"

	case 614:
		message = "invalid_grant"

	case 615:
		message = "unauthorized_client"

	case 616:
		message = "unsupported_grant_type"

	case 617:
		message = "invalid_scope"

	case 620:
		message = "角色名稱和標籤均為必填"

	case 621:
		message = "角色名稱或標籤重複"

	case 622:
		message = "Email 為必填"

	case 623:
		message = "角色不存在"

	case 624:
		message = "使用者不存在"

	case 625:
		message = "使用者已存在角色"

	case 626:
		message = "傳參錯誤，請檢查參數型態"

	case 627:  // 系統端
		message = "資料不可為空"

	case 702:
		message = "請求失敗，請聯絡客服"

	case 703:
		message = "系統別不吻合"

	case 704:
		message = "org_code 無資料"

	case 705:
		message = "該使用者不隸屬系統別"

	case 706:
		message = "Label 格式不符合規範"

	case 707:
		message = "Name 格式不符合規範"

	case 709:
		message = "存取拒絕: 查詢組織超出單位存取範圍"

	case 710:
		message = "存取拒絕: 單位存取範圍尚未設定"

	case 711:
		message = "查無資料"

	case 712:
		message = "請求參數為空"

	case 713:
		message = "未填寫有效參數"

	case 714:
		message = "傳參錯誤，請檢查參數型態"

	case 715:
		message = "per_page 上限 300"

	case 806:
		message = "更新失敗"

	case 807:
		message = "未填寫有效參數"

	case 808:
		message = "請填入符合系統之角色"

	case 809:
		message = "未輸入角色 id"

	case 810:
		message = "未輸入 google_user_id"

	case 811:
		message = "該使用者不隸屬系統別"

	case 812:
		message = "只能輸入一個角色"

	case 813:
		message = "Label 格式不符合規範"

	case 814:
		message = "Name 格式不符合規範"

	case 815:
		message = "角色名稱或標籤重複"

	case 816:
		message = "傳參錯誤，請檢查參數型態"

	case 901:
		message = "未填寫有效參數"

	case 902:
		message = "查無此角色"

	case 903:
		message = "傳參錯誤，請檢查參數型態"

	case 904:
		message = "該角色已綁定使用者，無法刪除"

	default:
		if env.IsLocal() {
			panic(fmt.Errorf("response message is not found for status code: %d !", statusCode))
		}
		message = "" // 無錯誤訊息

	}
	return message
}
