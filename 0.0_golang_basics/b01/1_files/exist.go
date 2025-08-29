package main

import (
	"fmt"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %v file1 ...\n", os.Args[0])
	}
	for _, file := range os.Args[1:] {
		stat, err := os.Stat(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on %v: %v\n", file, err)
			continue
		}
		if stat.IsDir() {
			fmt.Printf("%v is a directory\n", file)
		} else {
			fmt.Printf("%v is a file and has a size of %v bytes\n", file, stat.Size())
		}
		fmt.Printf("\tPerm:%v\n", stat.Mode().Perm())
	}
}
