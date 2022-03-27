package main

import "fmt"

/*
Two children, Lily and Ron, want to share a chocolate bar. Each of the squares has an integer on it.

Lily decides to share a contiguous segment of the bar selected such that:

The length of the segment matches Ron's birth month, and,
The sum of the integers on the squares is equal to his birthday.
Determine how many ways she can divide the chocolate.
*/
func main() {
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 5, 2))
}
func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	a := int32(0)
	count := int32(0)
	for a <= int32(len(s))-m {
		sum := int32(0)
		for i := a; i < a+m; i++ {
			sum += s[i]
		}
		if sum == d {
			count++
		}
		a++
	}
	return count

}
