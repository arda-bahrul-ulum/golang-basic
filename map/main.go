package main

import "fmt"

func main()  {
	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["golang"] = 9
	// myMap["js"] = 9
	// myMap["php"] = 8
	// myMap["nodejs"] = 9

	// fmt.Println(myMap["golang"])

	myMap := map[string]string{
		"golang": "is super fast", 
		"ruby": "is beautifull",
		"javscript": "hype",
	}

	// for key, value := range myMap {
	// 	fmt.Println("key : ", key, "value : ", value)
	// }

	// fmt.Println("====================================")

	// delete(myMap, "golang")

	// for key, value := range myMap {
	// 	fmt.Println("key : ", key, "value : ", value)
	// }

	value, isAvailable := myMap["php"]

	fmt.Println(value)
	fmt.Println(isAvailable)
}