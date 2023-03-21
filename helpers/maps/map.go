package maps

import (
	"encoding/json"
)

type m = map[string]interface{}

// Merge 合併兩 Map, 重複的 key left 優先
func Merge(left m, right m) map[string]interface{} {
	for key, rightVal := range right {
		// 左邊不存在的 key 以右邊補上
		if _, present := left[key]; !present {
			left[key] = rightVal
		}
	}
	return left
}
// FilterKeys 過濾 arg 只留下傳入的  allows 的 key 值
func FilterKeys(arg m, allows []string) map[string]interface{} {
	var ret = make(m)
	for _, attribute := range allows {
		if val, ok := arg[attribute]; ok {
			ret[attribute] = val
		}
	}
	return ret
}

// ToMap transform any object which may implement json tags to map
func ToMap(obj interface{}) (map[string]interface{}, error) {
	var err error
	objJson, err := json.Marshal(obj)
	if err != nil {
		return m{}, err
	}

	var result = make(m)
	err = json.Unmarshal(objJson, &result)

	return result, err
}

// ToMapSlice transform slice of objects to slice of maps
func ToMapSlice(obj interface{}) ([]m, error) {
	var err error
	objJson, err := json.Marshal(obj)
	if err != nil {
		return []m{}, err
	}

	var result []m
	err = json.Unmarshal(objJson, &result)

	return result, err
}

// ToPrettyJsonString 將 map 轉換為排版後的 json 字串
func ToPrettyJsonString(v m) (string, error)  {
	j, err := json.MarshalIndent(v, "", "  ")
	if  err != nil {
		return "", err
	}

	return string(j), nil
}

// ObjToPrettyJsonString 將 object 轉換為排版後的 json 字串
func ObjToPrettyJsonString(obj interface{})(string, error)  {
	if m, err := ToMap(obj); err != nil {
		return "", err
	} else if j, err := ToPrettyJsonString(m); err != nil {
		return "", err
	} else {
		return j, nil
	}
}
