package main

import (
	"fmt"
	"strings"
)

/*
You have two strings of lowercase English letters. You can perform two types of operations on the first string:

Append a lowercase English letter to the end of the string.
Delete the last character of the string. Performing this operation on an empty string results in an empty string.
Given an integer, k, and two strings, s and t, determine whether or not you can convert s to t by performing exactly
k of the above operations on s. If it's possible, print Yes. Otherwise, print No.
*/
func main() {
	x := appendAndDelete("ashley", "ashley", 10)
	fmt.Println(x)
}
func appendAndDelete(s string, t string, k int) string {
	splittedS := strings.Split(s, "")
	splittedT := strings.Split(t, "")
	var count int
	for i := 0; i == count && i < len(t); i++ {
		if splittedS[i] == splittedT[i] {
			count++
			continue
		}
	}
	rem := (len(s) + len(t) - (2 * count))

	if k >= rem && rem%2 == k%2 {
		return "Yes"
	} else if (len(s) + len(t) - int(k)) <= 0 {
		return "Yes"
	} else {
		return "No"
	}
	return "Yes"
}
