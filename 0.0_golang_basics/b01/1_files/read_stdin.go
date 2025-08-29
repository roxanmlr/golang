package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var buffer []byte = make([]byte, 0)
	if len(os.Args) != 1 {
		fmt.Fprintf(os.Stderr,"Usage: %v \n", os.Args[0])
		return
	}
	cumul := 0
	for  {
		var small_buf []byte = make([]byte, 1000)
		n, err := os.Stdin.Read(small_buf)
		if n == 0 {
			break;
		}
		if err != nil {
			log.Println(err)
			break;
		}
		buffer = append(buffer,small_buf[:n]...)
		cumul = cumul + n
	}
	os.Stdout.Write(buffer[:cumul])
}
