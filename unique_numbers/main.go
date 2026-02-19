package main

import (
	"fmt"
)

func main() {
	nums2 := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(numsUnique(nums2))
}

func numsUnique(nums []int) []int {
	numsMap := make(map[int]int)
	result := []int{}
	for _, num := range nums {
		numsMap[num]++
	}
	for num := range numsMap {
		result = append(result, num)
	}
	return result
}
