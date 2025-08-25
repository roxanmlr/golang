package main

import "os"

func ft_print_alphabet() {
	for i := range 26 {
		os.Stdout.Write([]byte{'a' + byte(i)})
	}
}

func main() {
	ft_print_alphabet()
}
