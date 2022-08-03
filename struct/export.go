package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func jsonify(p Person) string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("unable marshal to JSON")
	}
	return string(b)
}
func main() {

	p := Person{
		Name:  "Aлекс",
		Email: "alex@yandex.ru",
	}

	fmt.Println(jsonify(p))
}
