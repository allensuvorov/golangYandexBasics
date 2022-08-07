package main

import "fmt"

func main() {
	prices := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500}
	fmt.Println("перечень деликатесов — продуктов дороже 500 рублей:")
	for item, price := range prices {
		if price > 500 {
			fmt.Printf("%s - %d \n", item, price)
		}
	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var sum int
	for _, item := range order {
		sum += prices[item]
	}
	fmt.Printf("Total price for order %v is %v \n", order, sum)

}
