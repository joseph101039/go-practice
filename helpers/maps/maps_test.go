package maps

import (
	"fmt"
	"testing"
)

func Test_FilterKeys(t *testing.T) {

	m := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": struct{ a int }{1},
	}

	allows := [...]string{"c", "b"} // array, not a slice

	fmt.Println(
		FilterKeys(m, allows[:]), // 要將固定長度陣列轉換為 slice，可以使用切片運算符 [:]
	)

}

/**

請寫出一個函式, 輸入為一陣列lst及常數n, 輸出將lst分成每n個為一組的二維陣列,
example:
Input:  lst = [1, 2, 3, 4, 5], n=2
output:  result = [[1, 2], [3, 4], [5]]
*/

func Test_groupArray(t *testing.T) {
	lst := []int{1, 2, 3, 4, 5}
	n := 2
	fmt.Println(groupArray(lst, n))
}

func groupArray(lst []int, n int) (ret [][]int) {
	lstLen := len(lst)
	ret = [][]int{}
	for i := 0; i < lstLen; i += n {
		end := i + n
		if end > lstLen {
			end = lstLen
		}
		ret = append(ret, lst[i:end])
	}
	return
}

/**
請寫出一個函式, 輸入為一整數陣列lst, 找出符合條件的子數組數量,
*子數組條件為”數組中不同整數的個數剛好為k”
   *子數組內容相同,位置不同視為不同子數組,如下
example:
Input:  lst = [1, 3, 1, 3, 2], k=2
*符合條件的子數組 [1,3], [3,1], [1,3], [3,2], [1,3,1], [3,1,3], [1,3,1,3]
output:  result = 7
*/

func Test_arrange(t *testing.T) {
	lst := []int{1, 3, 1, 3, 2}
	k := 2
	fmt.Println(
		arrange(lst, k),
	)

}

func arrange(lst []int, k int) (result int) {

	dict := make(map[int]int)
	var start, end, unique int = 0, 0, 0

	for end < len(lst) {
		element := lst[end]
		if _, ok := dict[element]; !ok { // a new element

			for unique == k && start < end {
				dict[start]--
				if dict[start] == 0 {
					delete(dict, start)
					unique--
				}
				start++
			}

			unique++
			dict[element] = 1
		} else {
			dict[element]++
		}

		if unique == k {
			result++
		}
		end++
	}
	return
}
