package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var msg string
	//go operator(msg)
	//for {
	//	s := bufio.NewReader(os.Stdin)
	//	message, err := s.ReadString('\n')
	//	if err != nil {
	//		log.Fatalf("error reading input")
	//	}
	fmt.Scan(&msg)
	fmt.Println(msg)
	//operator(msg)
}

func operator(msg string) {
	var elements []int
	var result []int
	for _, val := range msg {
		str := strings.Split(string(val), " ")
		vall, _ := strconv.Atoi(str[1])
		if str[0] == "1" {
			elements = append(elements, vall)
		} else if str[0] == "2" {
			elements = elements[1:]
		} else if str[0] == "3" {
			first := elements[0]
			result = append(result, first)
			fmt.Println(string(first))
		}
	}
}
