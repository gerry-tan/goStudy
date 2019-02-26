package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	fmt.Printf("%v\n", stus)
	for i, stu := range stus {
		m[stu.Name] = &stus[i]
	}

	for k, v := range m {
		fmt.Printf("%s, %v\n", k, v)
	}
}
