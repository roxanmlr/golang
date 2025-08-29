package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
)

func main(){
	var str string
	var file *os.File
	var err error
	if !(len(os.Args) == 3 || 
		(len(os.Args) == 4 && strings.Compare(os.Args[1],"-v") == 0)){
		fmt.Fprintf(os.Stderr,"Usage: %v [-v] str file\n", os.Args[0])
		return
	}
	inverse := false
	if len(os.Args) == 4{
		inverse = true
	}
	if !inverse {
		str = os.Args[1]
		file, err = os.Open(os.Args[2])
	} else {
		str = os.Args[2]
		file, err = os.Open(os.Args[3])
	}
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if !inverse && strings.Contains(line, str){
			fmt.Print(line)
		}
		if inverse && !strings.Contains(line, str){
			fmt.Print(line)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
	}
}
