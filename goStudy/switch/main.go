package main

import "fmt"

func testSwitch1(number int) {
	num := number
	switch num {
	case 1:
		fmt.Println("num = 1")
	case 2, 3, 4:
		fmt.Println("num > 1 and num <= 4")
	case 5, 6, 7, 8:
		fmt.Println("num >=5 and num <= 8")
		fmt.Printf("current num is %d", num)
		fallthrough //穿透执行下个case
	case 9, 10:
		fmt.Println("num >= 9 and num <= 10")
	default:
		fmt.Println("num < 1 and num > 10")
	}
	fmt.Printf("current num: %d", num)
}

func testSwitch3(number int) {
	// num 作用域在 switch 中
	switch num := number; num {
	case 1:
		fmt.Println("num = 1")
	case 2, 3, 4:
		fmt.Println("num > 1 and num <= 4")
	case 5, 6, 7, 8:
		fmt.Println("num >=5 and num <= 8")
		fmt.Printf("current num is %d", num)
		fallthrough //穿透执行下个case
	case 9, 10:
		fmt.Println("num >= 9 and num <= 10")
	default:
		fmt.Println("num < 1 and num > 10")
	}
}

func testMulti() {
	// 99乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

func main() {
	//testSwitch1(1)
	//testSwitch3(5)
	testMulti()
}
