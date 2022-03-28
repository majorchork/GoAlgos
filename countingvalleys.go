package main

import "fmt"

/*
An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly  steps,
for every step it was noted if it was an uphill, , or a downhill,  step.
Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude.
We define the following terms:

A mountain is a sequence of consecutive steps above sea level,
starting with a step up from sea level and ending with a step down to sea level.
A valley is a sequence of consecutive steps below sea level,
starting with a step down from sea level and ending with a step up to sea level.
Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.
*/

func main() {
	fmt.Println(countingValleys(8, "UUDDDDDUU"))
	fmt.Println(countingValleys(10, "UUDDDUUDDDU"))
	fmt.Println(countingValleys(12, "UDDUDDDDUUUU"))
}

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	count := int32(0)
	valleys := int32(0)
	for _, val := range path {
		step := string(val)
		if step == "U" {
			count++
		} else {
			count--
		}
		if count == -1 && step != "U" {
			valleys++
		}
	}
	return valleys
}
