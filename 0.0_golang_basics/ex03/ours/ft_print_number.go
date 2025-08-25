package main

import "os"

func ft_print_number() {
	for i := range 10 {
		os.Stdout.Write([]byte{'0' + byte(i)})
	}
}

func main() {
	ft_print_number()
}
