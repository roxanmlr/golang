package main

import "os"

func ft_is_negative(c int) {
	if c < 0 {
		os.Stdout.Write([]byte{'N'})
	} else {
		os.Stdout.Write([]byte{'P'})
	}
}

func main() {
	ft_is_negative(15)
	ft_is_negative(-4)
}
