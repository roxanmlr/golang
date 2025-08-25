package main

import "os"

func ft_putchar(c byte) {
	os.Stdout.Write([]byte{c})
}

func main() {
	ft_putchar('a')
}
