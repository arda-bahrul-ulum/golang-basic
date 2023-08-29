package main

import "fmt"

func main()  {
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Andra", "score": "B"},
		{"name": "Bayu", "score": "C"},
	}	

	for _, student := range students {
		fmt.Println(student["name"], "-", student["score"])
	}

	fmt.Println(students)
}