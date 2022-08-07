package main

import "fmt"

type Item struct {
	NoOption   string
	Parameter1 string
	Parameter2 int
}

// funcops approach
type option func(*Item)

// constructor with options
func NewItem(opts ...option) *Item {
	i := &Item{
		NoOption:   "usual",
		Parameter1: "default",
		Parameter2: 42,
	}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

// funcopts

func Option1(option1 string) option {
	return func(i *Item) {
		i.Parameter1 = option1
	}
}

func Option2(option2 int) option {
	return func(i *Item) {
		i.Parameter2 = option2
	}
}

// initializing the object
func main() {
	// with default values
	item1 := NewItem()
	// with one option
	item2 := NewItem(Option2(70))
	// or with two options
	item3 := NewItem(Option1("unusual"), Option2(99))
	// options with changed order
	item4 := NewItem(Option2(88), Option1("rare"))
	fmt.Println(item1, item2, item3, item4)
}
