package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %v a b", os.Args[0])
	}
	max, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	for mul := range n {
		fmt.Printf("Table de %v\n", mul+1)
		print_table(mul+1, max)
		fmt.Println()
	}
}

func print_table(mul int, max int) {
	for i := range max + 1 {
		fmt.Printf("%v x %v = %v\n", mul, i, mul*i)
	}
}
