package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "unknown"
	if len(os.Args) >= 2 {
		name = os.Args[1]
	}
	age := "unknown"
	if len(os.Args) >= 3 {
		age = os.Args[2]
	}
	fmt.Printf("Hello %v, You are %v years old", name, age)
}
