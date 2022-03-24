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
func isValid(s string) bool {

	var braces []string
	m := map[string]string{"[": "]", "(": ")", "{": "}", "<": ">"}
	for _, i := range s {
		bc := string(i)
		_, ok := m[bc]
		if ok {
			braces = append(braces, bc)
		} else if len(braces) == 0 {
			return false
		} else {
			last := braces[len(braces)-1]
			if m[last] != bc {
				return false
			} else {
				braces = braces[:len(braces)-1]
				// continue
			}
		}
	}
	if len(braces) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("({){}}{{}{{]{))([[][]][)({()}{"))
	fmt.Println(isValid("{}()[]"))
	fmt.Println(ValidBraces("[]{}()()()[][][]{}{}[]{}[]())("))
	fmt.Println(ValidBraces("{(["))
	fmt.Println(ValidBraces("{()}"))
	fmt.Println(ValidBraces("({){}}{{))([[][]][)({()}{"))
	fmt.Println(ValidBraces("[]{}()()()[][][]{}{}[]{}[]()"))

}
