package main

import "fmt"

var Global = 5

func Restore(n *int) *int {
	*n--
	fmt.Println("restored=", *n)
	return n
}

func main() {
	Global++
	fmt.Println(Global)
	defer Restore(&Global)
}
