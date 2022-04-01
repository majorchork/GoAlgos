package main

import "fmt"

func main() {
	fmt.Println(boomerang([]int{5, 9, 5, 9, 5}))
	fmt.Println(boomerang([]int{9, 5, 9, 5, 1, 1, 1}))
	fmt.Println(boomerang([]int{5, 6, 6, 7, 6, 3, 9}))
	fmt.Println(boomerang([]int{3, 7, 3, 2, 1, 5, 1, 2, 2, -2, 2}))
	fmt.Println(boomerang([]int{2, 2, 2, 2, 3, 2, 3}))
	fmt.Println(boomerang([]int{3, 7, 3, 2, 1, 5, 1, 2, 2, -2, 2}))
	//3, 7, 3, 2, 1, 5, 1, 2, 2, -2, 2
}
func boomerang(boo []int) int {
	var count int
	for i := 0; i < len(boo)-2; i++ {
		var k = i + 1
		var j = i + 2
		if boo[i] == boo[j] && boo[k] != boo[i] {
			count++
		}
	}
	return count
}
