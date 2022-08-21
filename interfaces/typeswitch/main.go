package main

import (
	"fmt"
	"strings"
)

func Mul(a interface{}, b int) interface{} {
	// switch v := a.(type) {
	// case string:
	// 	return func() {
	// 		for i := 1; i <= b; i++ {
	// 			v += v
	// 		}
	// 		fmt.Println(v)
	// 	}()
	// case int:
	// 	fmt.Println(v * b)
	// default:
	// 	return "unknown type"
	// 	// fmt.Println("unknown type")
	// }
	switch va := a.(type) {
	case int:
		return va * b
	case string:
		return strings.Repeat(va, b)
	case fmt.Stringer:
		return strings.Repeat(va.String(), b)
	default:
		return nil
	}
}

func main() {
	fmt.Println(Mul("king", 2))
}
