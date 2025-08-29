package main

import (
	"fmt"
	"os"
	"regexp"
)

func main(){
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr,"Usage: %v str ...\n", os.Args[0])
		return
	}
	var validList = regexp.MustCompile(`^\s*\[\s*"([^"]*)"\s*(,\s*"([^"]*)"\s*)*?\]\s*$`)
	var captureString = regexp.MustCompile(`"([^"])"`)
	for _, str := range(os.Args[1:]) {
		fmt.Printf("%v match %v\n", str, validList.MatchString(str))
		s := captureString.FindAllStringSubmatch(str, -1)
		for _, t := range s{
			fmt.Printf("\t%v\n", t[1])
		}
	}
}
