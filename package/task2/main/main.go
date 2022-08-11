package main

import (
	"fmt"

	"github.com/allensuvorov/golangYandexBasics/tree/main/package/task2/mypack"
	//mypack "github.com/allensuvorov/golangYandexBasics/tree/main/package/task2/mypack"
)

func main() {
	if sum := mypack.Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
