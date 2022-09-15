package task4

import "fmt"

func EstimateValue(value int) string {
	switch {
	case value < 10:
		return "small"

	case value < 100:
		return "medium"

	default:
		return "big"
	}
}

func main() {
	v := 10
	fmt.Println(EstimateValue(v))
}
