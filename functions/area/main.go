package main

type figures int

const(
    square figures = iota // квадрат
    circle // круг
    triangle // равносторонний треугольник
)

func init() {
	areas := make(map[figures]func(figures)float64)
}

func area(f figures)(func(float64) float64, bool) {
	ok := f 
	ar := func(float64) float64 {

	}
	return ar, ok
}

ar, ok := area(myFigure)
if !ok {
    fmt.Println("Ошибка")
    return
}
myArea := ar(x)

func main() {

}