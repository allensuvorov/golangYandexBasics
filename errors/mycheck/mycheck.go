package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck
//
// ...

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
