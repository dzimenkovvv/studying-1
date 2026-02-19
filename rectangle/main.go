package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func main() {
	rectangle := Rectangle{
		width:  10,
		height: 3,
	}
	fmt.Println("Площадь:", rectangle.Area())
	fmt.Println("Периметр:", rectangle.Perimeter())
}

func (r Rectangle) Area() int {
	return r.height * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.height + r.width)
}
