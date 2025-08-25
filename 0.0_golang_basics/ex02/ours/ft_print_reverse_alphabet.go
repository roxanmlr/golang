package main

import "os"

func ft_print_reverse_alphabet() {
	for i := range 26 {
		os.Stdout.Write([]byte{'z' - byte(i)})
	}
}

func main() {
	ft_print_reverse_alphabet()
}
