package main

import (
	"fmt"
	"io"
	"os"
)


func main(){
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,"Usage: %v src dest")
		return
	}
	src, err := os.Open(os.Args[1])	
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n", err)
		return
	}
	defer src.Close()
	dest, err := os.Create(os.Args[2])	
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n", err)
		return
	}
	defer dest.Close()
	io.Copy(dest, src)
}	
