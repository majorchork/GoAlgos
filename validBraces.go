package main

import "fmt"

func ValidBraces(braces string) bool {
	openBraces := []string{}

	for _, eachBrace := range braces {
		s := string(eachBrace)

		if s == "(" || s == "{" || s == "[" {
			openBraces = append(openBraces, s)
		} else {
			if len(openBraces) == 0 {
				return false
			}
			lastBrace := openBraces[len(openBraces)-1]
			if s == ")" && lastBrace == "(" {
				openBraces = openBraces[:len(openBraces)-1]
				continue
			} else if s == "}" && lastBrace == "{" {
				openBraces = openBraces[:len(openBraces)-1]
				continue
			} else if s == "]" && lastBrace == "[" {
				openBraces = openBraces[:len(openBraces)-1]
				continue
			}
			return false
		}
	}

	if len(openBraces) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(ValidBraces("({){}}{{}{{]{))([[][]][)({()}{"))
	fmt.Println(ValidBraces("{}()[]"))
	fmt.Println(ValidBraces("[]{}()()()[][][]{}{}[]{}[]())("))
	fmt.Println(ValidBraces("{(["))
	fmt.Println(ValidBraces("{()}"))
	fmt.Println(ValidBraces("({){}}{{))([[][]][)({()}{"))
	fmt.Println(ValidBraces("[]{}()()()[][][]{}{}[]{}[]()"))

}
