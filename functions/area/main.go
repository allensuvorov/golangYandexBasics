package main

import (
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func init() {
	areas := make(map[figures]func(float64) float64)
	areas[square] = func(x float64) float64 {
		return x * x
	}
	areas[circle] = func(x float64) float64 {
		return x * x * math.Pi

	}
	areas[triangle] = func(x float64) float64 {
		return x * x * math.Sqrt(3) / 4
	}

}

// func area(f figures)(func(float64) float64, bool) {
// 	ok := f
// 	ar := func(float64) float64 {

// 	}
// 	return ar, ok
// }

// ar, ok := area(myFigure)
// if !ok {
//     fmt.Println("Ошибка")
//     return
// }
// myArea := ar(x)

// func main() {

// }
