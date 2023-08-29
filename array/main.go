package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "GO"
	// languages[1] = "JS" 
	// languages[2] = "REACT" 
	// languages[3] = "VUE" 
	// languages[4] = "NODE" 

	// languages := [5]string{
	// 	"a", 
	// 	"b", 
	// 	"c", 
	// 	"d", 
	// 	"e",
	// }

	languages := [...]string{
		"a", 
		"b", 
		"c", 
		"d", 
		"e",
		"f",
	}

	for index, lang := range languages {
		fmt.Println("Index : ", index, "languages : ", lang)
	}	

	// fmt.Println(languages)

	// length := len(languages)

	// fmt.Println(length)
}