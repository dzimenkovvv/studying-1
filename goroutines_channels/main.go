package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan int)

	go func() {
		iterations := 1 + rand.Intn(10)
		fmt.Println("Количество итераций:", iterations)
		for i := 1; i <= iterations; i++ {
			ch <- 1 + rand.Intn(5)
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
