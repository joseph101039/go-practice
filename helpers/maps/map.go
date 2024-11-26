package maps

import (
	"encoding/json"
)

type m = map[string]any

// Merge 合併兩 Map, 重複的 key left 優先
func Merge(left, right m) map[string]any {
	for key, rightVal := range right {
		// 左邊不存在的 key 以右邊補上
		if _, present := left[key]; !present {
			left[key] = rightVal
		}
	}
	return left
}

// FilterKeys 過濾 arg 只留下傳入的  allows 的 key 值
func FilterKeys[T comparable](arg map[T]any, allows []T) map[T]any {
	var ret = make(map[T]any)
	for _, attribute := range allows {
		if val, ok := arg[attribute]; ok {
			ret[attribute] = val
		}
	}

	return ret
}

// ToMap transform any object which may implement json tags to map
func ToMap(obj any) (map[string]any, error) {
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
func ToMapSlice(obj any) ([]m, error) {
	var err error
	objJson, err := json.Marshal(obj)
	if err != nil {
		return []m{}, err
	}

	var result []m
	err = json.Unmarshal(objJson, &result)

	return result, err
}

// ObjToPrettyJsonString 將轉換為排版後的 json 字串
func ToJson(v any) (j string, err error) {
	var o any
	switch s := v.(type) {
	case string:
		if err = json.Unmarshal([]byte(s), &o); err != nil {
			o = s
		}

	case []byte:
		if err = json.Unmarshal(s, &o); err != nil {
			o = s
		}
	default:
		o = s
	}

	var b []byte
	b, err = json.MarshalIndent(o, "", "  ")
	j = string(b)
	return
}

/**
請寫出一個函式, 輸入為一陣列lst及常數n, 輸出將lst分成每n個為一組的二維陣列,
example:
Input:  lst = [1, 2, 3, 4, 5], n=2
output:  result = [[1, 2], [3, 4], [5]]
*/

func rerange[T []any](lst []T, n int) (ret [][]T) {
	lstLen := len(lst) / n
	ret = make([][]T, lstLen)
	for i := 0; i < lstLen; i++ {
		arr := make([]T, n)
		for j := 0; j < n; j++ {
			arr[j] = lst[i*n+j]
		}
		ret[i] = arr
	}

	return
}
