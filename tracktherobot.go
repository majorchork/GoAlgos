package main

import "fmt"

/*
This robot roams around a 2D grid. It starts at (0, 0) facing North.
After each time it moves, the robot rotates 90 degrees clockwise.
Given the amount the robot has moved each time,
you have to calculate the robot's final position.
NOTE THAT everytime the robot moves north you increment the y aix by distance moved
each time the robot moved east you increment in the x axis
each time it moved south you decrease in the y axis
each time it moved west you decrement in the x axis
*/

func main() {
	fmt.Println(trackTheRobot([]int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0}))
	fmt.Println(trackTheRobot([]int{}))
	fmt.Println(trackTheRobot([]int{10, -10, -10, 10}))
	fmt.Println(trackTheRobot([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(trackTheRobot([]int{1, 0, 2, 0, 3, 0, 4, 0, 5, 0}))
	fmt.Println(trackTheRobot([]int{0, 1, 0, 2, 0, 3, 0, 4, 0, 5}))
}
func trackTheRobot(dir []int) []int {
	var result []int
	var x int
	var y int
	var a int = 1
	for _, v := range dir {
		if a == 1 {
			y += v
			a++
		} else if a == 2 {
			x += v
			a++
		} else if a == 3 {
			y = y - v
			a++
		} else if a == 4 {
			x = x - v
			a = 1
		}
	}
	result = append(result, x)
	result = append(result, y)
	return result
}
