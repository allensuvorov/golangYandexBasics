package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	// fmt.Println(input[6])
	fmt.Println(RemoveDuplicates(input))
}

// remove duplicate while keeping the order in the slice

func RemoveDuplicates(input []string) []string {
	m := make(map[string]int) // uniqs
	output := input
	for i, v := range output {
		if _, ok := m[v]; ok { // delete from slice
			if i == len(output)-1 {
				output = output[:i]
				break
			} else {
				output = append(output[:i], output[i+1:]...)
			}
		} else {
			m[v] = i
		}
	}
	return output
}
