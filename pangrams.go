package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("abcdefghijklmnopqrstuvwxyz"))
}

func pangrams(s string) string {
	var container []string
	newS := strings.ToLower(s)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range alphabet {
		for _, value := range newS {
			if char == value {
				container = append(container, string(char))
				break
			}
		}
		if len(container) == 26 {
			return "pangram"
		}
	}
	return "not pangram"
}
