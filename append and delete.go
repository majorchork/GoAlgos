package main

import (
	"fmt"
	"strings"
)

func main() {
	x := appendAndDelete("ashley", "ash", 2)
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
	tRem := len(t) - count
	sRem := len(s) - count
	sTK := k - (tRem + sRem)
	if count == len(t) && len(t) == len(s) && int(k) > count {
		return "Yes"
	}
	if tRem+sRem < k && sTK%2 != 0 {
		return "No"
	} else if (tRem + sRem) <= k {
		return "Yes"
	} else {
		return "No"
	}
}
