package main

import (
	"first-project/calculation"
	"fmt"
)

func main() {
	fmt.Println("hello world!")
	// sentence := TestAja()

	// fmt.Println(sentence)

	result := calculation.Add(10,9)

	fmt.Println(result)
}