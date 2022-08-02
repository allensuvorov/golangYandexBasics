package main

import "fmt"

func main() {
	nums := [5]int{11, 22, 33, 44, 55}
	k := 66
	var n1, n2 int

	// put array in map
	m := make(map[int]int, 5)
	for i, num := range nums {
		m[num] = i
	}
	// find the indices
	for num, i := range m {
		if _, ok := m[k-num]; ok {
			n1, n2 = i, m[k-num]
			break
		}
	}
	fmt.Println(n1, n2)
}
