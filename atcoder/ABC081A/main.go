package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Scan(&input)

	result := 0
	for _, c := range input {
		if string(c) == "1" {
			result++
		}
	}

	fmt.Println(result)
}
