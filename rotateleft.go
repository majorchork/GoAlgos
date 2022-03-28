package main

import "fmt"

func main() {
	fmt.Println(rotateLeft(3, []int32{1, 2, 3, 4, 5}))
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
