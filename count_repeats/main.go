package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 2, 3, 3, 3}
	fmt.Println(numsCount(nums1))
}
func numsCount(nums []int) map[int]int {
	result := make(map[int]int)
	for _, num := range nums {
		result[num]++
	}
	return result
}
