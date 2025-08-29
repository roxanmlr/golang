package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr,"Usage: %v op1 ...\n", os.Args[0])
		return
	}
	for _, str := range os.Args[1:] {
		ops := strings.Split(str, " ")
		if len(ops) != 3 && len(ops) != 1 {
			fmt.Fprintf(os.Stderr,"\"%v\" not correctly formatted. Put spaces between operandes.\n", str)
			continue
		}
		nbr1, err := strconv.ParseFloat(ops[0], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			continue
		}
		if len(ops) == 1 {
			fmt.Printf("%v = %v\n", str, nbr1)
			continue
		}
		nbr2, err := strconv.ParseFloat(ops[2], 32)
		if err != nil {
			fmt.Fprintf(os.Stderr,"%v\n", err)
			continue
		}
		var res float64
		var fail bool = false
		switch(ops[1]){
			case "+":
				res = nbr1 + nbr2
			case "-":
				res = nbr1 - nbr2
			case "*":
				res = nbr1 * nbr2
			case "/":
				res = nbr1 / nbr2
			default:
				fmt.Fprintf(os.Stderr,"%v operand unknown\n")
				fail = true
		}
		if fail {
			continue
		}
		fmt.Printf("%v = %v\n", str, res)
	}
}
