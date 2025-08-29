package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr,"Usage: %v str ...\n", os.Args[0])
		return
	}
	for _, str := range(os.Args[1:]) {
		var arr []string
		err := json.Unmarshal([]byte(str), &arr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, a := range arr {
			fmt.Printf("\t\"%v\"\n", a)
		}
	}
}
