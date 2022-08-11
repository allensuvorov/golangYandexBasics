package main

import (
	"fmt"

	"yourpackage" // ваш пакет
)

func main() {
	if sum := yourpackage.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
