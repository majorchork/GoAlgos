package main

import "fmt"

func sumProduct(arr ...int) int {
	sum := 0
	// adds all elements
	for _, v := range arr {
		sum += v
	}
	return helper(sum)
}
func helper(sum int) int {
	// Base Case
	if sum < 10 {
		return sum
	}
	return helper(swap(sum))
}

func swap(n int) int {
	product := 1
	for n > 0 {

		lastDigit := n % 10  // 4 => 4 % 10 = 4
		product *= lastDigit // 4 * 4 = 16
		n = n / 10           // 4/10 => 0
	}
	return product
}

func main() {
	fmt.Println(sumProduct(98526, 54, 863, 156489, 45, 6156))
}
