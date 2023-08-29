package main

import "fmt"

func main()  {
	// sentence := printMyResult("saya sedang")
	// fmt.Println(sentence)

	result := add(10, 100)
	fmt.Println(result)
}

// func printMyResult(sentence string) string {
// 	newSentence := sentence + " belajar golang!"
// 	return newSentence
// }

func add(number, numberTwo int) int {
	return number + numberTwo
}

// input
// proccess
// output