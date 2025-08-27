package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage:%v string_to_print times\n", os.Args[0])
		//os.Args[0] est le nom du programme
		return
	}
	n, _ := strconv.Atoi(os.Args[2])
	for i := range n {
		fmt.Println(i+1, ": ", os.Args[1])
	}
}
