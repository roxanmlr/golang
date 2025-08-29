package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strconv"
	"os"
)

type Person struct {
	Name	string	`json:"name"`
	Age		int		`json:"age"`
}

func main(){
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr,"Usage: %v str ...\n", os.Args[0])
		return
	}
	var l []Person = make([]Person, len(os.Args)-1)
	for i, str := range(os.Args[1:]) {
		err := json.Unmarshal([]byte(str), &(l[i]))
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			continue
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil && len(scanner.Text()) != 0{
			fmt.Fprintf(os.Stderr,"%v\n", err)
			fmt.Print(">> ")
			continue
		}
		if len(scanner.Text()) == 0 {
			break
		}
		if 0 > i || i >= int64(len(l)) {
			fmt.Fprintf(os.Stderr,"Can't reach %v indice\n", i)
			fmt.Print(">> ")
			continue
		}
		fmt.Printf("Person{Name: \"%v\",Age: %v}\n>> ", l[i].Name, l[i].Age) 
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
