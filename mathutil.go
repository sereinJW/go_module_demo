package mathutil

import (
	"fmt"
)

func Add(a, b int) int {
	c := a + b
	fmt.Println("这是两数相加")
	return c
}

func Multiply(a, b int) int {
	c := a * b
	fmt.Println("这是两数相乘")
	return c
}
