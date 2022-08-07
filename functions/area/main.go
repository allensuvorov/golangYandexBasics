package main

import (
	"fmt"
	"math"
)

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

var areas = make(map[figures]func(float64) float64)

func area(f figures) (func(float64) float64, bool) {
	areas[square] = func(x float64) float64 {
		return x * x
	}
	areas[circle] = func(x float64) float64 {
		return x * x * math.Pi

	}
	areas[triangle] = func(x float64) float64 {
		return x * x * math.Sqrt(3) / 4
	}
	ar, ok := areas[f]
	return ar, ok
}

func main() {
	ar, ok := area(square)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	myArea := ar(3)
	fmt.Println(myArea)
}
