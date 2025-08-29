package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var buffer []byte = make([]byte, 1000)
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,"Usage: %v src dest\n", os.Args[0])
		return
	}
	src, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	defer src.Close()
	dest, err := os.Create(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}
	defer dest.Close()
	for  {
		n, err := src.Read(buffer)
		if n == 0 {
			break;
		}
		if err != nil {
			log.Println(err)
			break;
		}
		dest.Write(buffer[:n])
	}
}
