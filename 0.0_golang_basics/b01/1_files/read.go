package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var buffer []byte = make([]byte, 1000)
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr,"Usage: %v file1 ...\n", os.Args[0])
		return
	}
	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			log.Println(err)
			continue
		}
		defer file.Close()
		for  {
			n, err := file.Read(buffer)
			if n == 0 {
				break;
			}
			if err != nil {
				log.Println(err)
				break;
			}
			os.Stdout.Write(buffer[:n])
		}
	}
}
