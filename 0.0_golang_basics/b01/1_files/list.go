package main

import (
	"fmt"
	"log"
	"os"
)

func isDir(name string)bool {
	stat, err := os.Stat(name)
	if err != nil {
		log.Println(err)
		return false
	}
	if stat.IsDir() {
		return true
	}
	return false
}

func main(){
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %v dir1 ...", os.Args[0])
	}
	for _, dir := range os.Args[1:] {
		if !isDir(dir){
			fmt.Println(dir)
			continue
		}
		dirEntries, err := os.ReadDir(dir)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(os.Args) > 2 {
			fmt.Printf("%v:\n", dir)
		}
		for _, d := range dirEntries {
			fmt.Println(d.Name())
		}
		fmt.Println()
	}
}
