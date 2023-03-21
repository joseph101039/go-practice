package main

func SortParallel(input []int, ch chan []int) {
	if len(input) < 2 {
		ch <- input
		return
	}

	mid := int(len(input) / 2)
	ch1, ch2 := make(chan []int), make(chan []int)
	go SortParallel(input[:mid], ch1)
	go SortParallel(input[mid:], ch2)
	ch <- mergesort(<-ch1, <-ch2)

}

func Sort(input []int) []int {
	if len(input) < 2 {
		return input
	}

	mid := int(len(input) / 2)
	return mergesort(
		Sort(input[:mid]),
		Sort(input[mid:]),
	)
}

func mergesort(a1 []int, a2 []int) []int {
	a := []int{}
	//fmt.Println(a1, a2, len(a1), len(a2))
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			a = append(a, a1[i])
			i++
		} else if a1[i] > a2[j] {
			a = append(a, a2[j])
			j++
		}

	}
	for i < len(a1) {
		a = append(a, a1[i])
		i++
	}
	for j < len(a2) {
		a = append(a, a2[j])
		j++
	}
	//fmt.Println("a is ", a)
	return a

}
