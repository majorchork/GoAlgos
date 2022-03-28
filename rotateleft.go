package main

import "fmt"

/*
A left rotation operation on an array of size n shifts each of the array's elements 1 unit to the left.
Given an integer, d, rotate the array that many steps left and return the result.

Function Description

rotateLeft has the following parameters:

int d: the amount to rotate by
int arr[n]: the array to rotate
*/

func main() {
	fmt.Println(rotateLeft(3, []int32{1, 2, 3, 4, 5}))
	fmt.Println(rotateLeft(4, []int32{1, 2, 3, 4, 5, 6}))
	fmt.Println(rotateLeft(3, []int32{1, 2, 3, 4, 5}))
	fmt.Println(rotateLeft(2, []int32{1, 2, 3, 4, 5}))
	fmt.Println(rotateLeft(1, []int32{1, 2, 3, 4, 5}))

}

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	i := int32(0)
	for i < d {
		first := arr[0]
		arr = arr[1:]
		arr = append(arr, first)
		i++
	}
	fmt.Println(arr)

	a := arr[:d]
	b := arr[d:]
	b = append(b, a...)
	return b
}
