package main

import "fmt"

func main() {
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "A")
	gamingConsole = append(gamingConsole, "B")
	gamingConsole = append(gamingConsole, "C")

	// gamingConsole := []string{"and the game world will interact with the game world", "and the game"}

	for _, console := range gamingConsole {
		fmt.Println(console)
	}
	// fmt.Println(gamingConsole)

}