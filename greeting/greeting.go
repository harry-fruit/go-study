package main

import (
	"fmt"
)

func main() {
	name := getName()
	fmt.Printf(name)

	// if name != "" {
	// 	fmt.Print("Hello, ", name)
	// 	return
	// }

	// fmt.Print("Hello")
}

func getName() string {
	var name string

	fmt.Print("What is your name?\n")
	_, err := fmt.Scanf("%s", &name) // Added error handling
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	return name
}
