package main

import "strconv"

func main() {
	getMax([]string{"1 11", "1, 24", "1 36", "2", "3", "1 45", "1 54", "2", "1 20", "3"})
}

func getMax(operations []string) []int32 {
	// Write your code here
	var resultArr []int
	var maxArr []int32

	for _, item := range operations {
		if item[0] == 49 {
			newVal, _ := strconv.Atoi(item[2:])
			resultArr = append(resultArr, newVal)
		} else if item[0] == 50 {
			resultArr = resultArr[:len(resultArr)-1]
		} else if item[0] == 51 {
			maxVal := int32(0)
			for _, val := range resultArr {
				if int32(val) > maxVal {
					maxVal = int32(val)
				}
			}
			maxArr = append(maxArr, maxVal)
		}
	}
	return maxArr
}
