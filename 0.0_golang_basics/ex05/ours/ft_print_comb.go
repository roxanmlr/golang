package main

import "os"

func ft_print_comb() {
	var num []byte = []byte("012")
	for num[0] <= '7' {
		num[1] = num[0] + 1
		for num[1] <= '8' {
			num[2] = num[1] + 1
			for num[2] <= '9' {
				os.Stdout.Write(num)
				if !(num[0] == '7' && num[1] == '8' && num[2] == '9') {
					os.Stdout.Write([]byte(", "))
				}
				num[2]++
			}
			num[1]++
		}
		num[0]++
	}
}

func main() {
	ft_print_comb()
}
