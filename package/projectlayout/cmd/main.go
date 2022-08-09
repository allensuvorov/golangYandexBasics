package main

import (
	"fmt"
	mathslice "yandex/golangYandexBasics/package/projectlayout/internal/"
)

func main() {

	s := mathslice.Slice{1, 2, 3}
	fmt.Println(s)
	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	mathslice.MapSlice(s, func(i mathslice.Element) mathslice.Element {
		return i * 2
	})

	fmt.Println("Слайс, умноженный на два: ", s)

	fmt.Println("Сумма слайса: ", mathslice.SumSlice(s))

	fmt.Println("Свёртка слайса умножением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x * y
			},
			1))

	fmt.Println("Свёртка слайса сложением ",
		mathslice.FoldSlice(s,
			func(x mathslice.Element, y mathslice.Element) mathslice.Element {
				return x + y
			},
			0))

}
