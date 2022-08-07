package main

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
}
