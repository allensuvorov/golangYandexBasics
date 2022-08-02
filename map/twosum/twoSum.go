package main

import "fmt"

func twoSum() []int {
	nums := [5]int{11, 22, 33, 44, 55}
	k := 99
	var n1, n2 int

	// look for second num and build map
	m := make(map[int]int, 5)
	for i, num := range nums {
		if _, ok := m[k-num]; ok {
			n1, n2 = i, m[k-num]
			fmt.Println(n1, n2)
			return []int{n1, n2}
		}
		m[num] = i
	}
	// fmt.Println(n1, n2)
	return nil
}
func main() {
	twoSum()
}
