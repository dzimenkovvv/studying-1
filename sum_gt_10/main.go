package main

import (
	"fmt"
)

func main() {
	var nums, sum int
	slice := []int{}
	for {
		fmt.Scan(&nums)
		if nums != -1 {
			slice = append(slice, nums)
			if nums > 10 {
				sum += nums
			}
		} else {
			break
		}
	}
	fmt.Println("Введенные данные:", slice)
	fmt.Println("Сумма введенных чисел > 10:", sum)
}
