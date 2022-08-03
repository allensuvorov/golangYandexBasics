package main

import "fmt"

func TwoSum(nums []int, k int) []int {
	var n1, n2 int
	// look for second num and build map
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if _, ok := m[k-num]; ok {
			n1, n2 = i, m[k-num]
			fmt.Println(n1, n2)
			return []int{n1, n2}
		}
		m[num] = i
	}
	return nil
}
func main() {
	nums := []int{11, 22, 33, 44, 55}
	k := 99
	TwoSum(nums, k)
}
