package request

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllQueryParams(context *gin.Context) map[string][]string {
	return context.Request.URL.Query()
}

// GetBodyParams 取得不同 Content-Type 的參數 (For POST, PATCH, PUT)
func GetBodyParams(context *gin.Context) map[string]interface{} {
	requestValue := make(map[string]interface{})
	header := context.GetHeader("Content-Type")
	if strings.HasPrefix(header, "application/json") { // 判斷開頭是 application/json
		context.BindJSON(&requestValue)
	} else {
		context.MultipartForm()
		for key, value := range context.Request.PostForm {
			if len(value) > 1 {
				requestValue[key] = value
			} else {
				requestValue[key] = value[0]
			}
		}
	}

	return requestValue
}

// GetValue 解析 bodyParameters 內簡單值, 若是解析 array, slice, map 會報錯
func GetValue(bodyParams map[string]interface{}, key string) string {
	if val, ok := bodyParams[key]; ok {
		return toString(val)
	} else {
		return ""
	}
}

func GetIValue(bodyParams map[string]interface{}, key string) interface{} {
	if val, ok := bodyParams[key]; ok {
		return val
	} else {
		return ""
	}
}

// DefaultGetValue 當 GetValue 為空字串時回傳 defaultValue
func DefaultGetValue(bodyParams map[string]interface{}, key string, defaultValue string) string {
	if val := GetValue(bodyParams, key); val == "" {
		return defaultValue
	} else {
		return val
	}
}

// toString formats simple value as a string. 不支援 array, slice, map 等格式
func toString(value interface{}) string {
	if strVal, err := formatBasicType(reflect.ValueOf(value)); err == nil {
		return strVal
	} else {
		// 傳入不支援的型態
		panic(err)
	}
}

func FormatString(v any) (val string, err error) {

	switch s := v.(type) {
	case string:
		val = s
	case []byte:
		val = string(s)
	case []rune:
		val = string(s)


	case int:
		val = strconv.FormatInt(int64(s), 10)
	case int8:
		val = strconv.FormatInt(int64(s), 10)
	case int16:
		val = strconv.FormatInt(int64(s), 10)
	case int32:
		val = strconv.FormatInt(int64(s), 10)
	case int64:
		val = strconv.FormatInt(int64(s), 10)
	case uint:
		val = strconv.FormatUint(uint64(s), 10)
	case uint8:
		val = strconv.FormatUint(uint64(s), 10)
	case uint16:
		val = strconv.FormatUint(uint64(s), 10)
	case uint32:
		val = strconv.FormatUint(uint64(s), 10)
	case uint64:
		val = strconv.FormatUint(uint64(s), 10)
	case uintptr:
		val = strconv.FormatUint(uint64(s), 10)
	case float32:
		val = strconv.FormatFloat(float64(s), 'f', -1, 64)
	case float64:
		val = strconv.FormatFloat(float64(s), 'f', -1, 64)
	case bool:
		if s {
			val = "1"
		} else {
			val = "0"
		}
	case nil:
		val = ""
	default:
		err = fmt.Errorf("cannot convert type %t to string", s)
	}

	return
}

// formatBasicType formats a value without inspecting its internal structure.
// reference: https://wizardforcel.gitbooks.io/gopl-zh/content/ch12/ch12-02.html
func formatBasicType(v reflect.Value) (string, error) {
	switch v.Kind() {
	case reflect.Invalid:
		return "", nil // 空值 (nil)
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil // 64 bit 無科學記號
	case reflect.Bool:
		if v.Bool() {
			return "1", nil
		} else {
			return "0", nil
		}
	case reflect.String:
		return v.String(), nil
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return "", fmt.Errorf("%s is not supported", v.Type().String())
	}
}
