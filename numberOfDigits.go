package main

import "fmt"

func numberOfDigits(n int) int {
	//Find Number of Digits in Number
	//Create a function that will return an integer number corresponding to the amount of digits in the given integer num.
	//	Examples
	//num_of_digits(1000) ➞ 4
	//num_of_digits(12) ➞ 2
	//num_of_digits(1305981031) ➞ 10
	//num_of_digits(0) ➞ 1

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
