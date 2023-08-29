package main

import "fmt"

func main()  {
	
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("learning golang! :", i)
	// }

	// i := 1;
	// for i <= 100 {
	// 	fmt.Println("learning golang! :", i)
	// 	i++
	// }

	title := "Golang the best language"

	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }

	for index, character := range title {
		if (index %2 == 0){
			fmt.Println("index :", index, "character :", string(character))
		}
	}

	for index, letter := range title {
		letterString := string(letter)

		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("index :", index, " letter :", string(letter))
		// }

		switch letterString {
			case "a", "i", "u", "e", "o":
				fmt.Println("index :", index, " letter :", string(letter));
		}
	}
}