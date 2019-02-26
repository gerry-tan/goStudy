package main

import "fmt"

//多返回值
func calc1() (int, int) {
	return 1, 3
}

//可变参数
func calc2(a int, b ...int) int {
	sum := a
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func main() {
	a, b := calc1()
	fmt.Printf("a=%d  b=%d\n", a, b)

	sum := calc2(1, 2, 3, 4)
	fmt.Printf("sum=%d\n", sum)
}
