package main

import "fmt"

func numberOfDigits(n int) int {
	var product []int
	if n == 0 {
		return 1
	}
	for n > 0 {
		lastDigit := n % 10 // 4 => 4 % 10 = 4
		product = append(product, lastDigit)
		n = n / 10
		continue
	}
	return len(product)
}
func main() {
	fmt.Println(numberOfDigits(1305955))
}
