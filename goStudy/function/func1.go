package main

import "fmt"

func GetValue(m map[int]string, id int) (str string, b bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return str, false
}

const (
	x = iota
	y
	z = "zz"
	m
	n = iota
)

func main() {
	var i = 20
	fmt.Println(x, y, z, m, n, &i)

	type MyInt1 int
	type MyInt2 = int
	var t = 9
	var i1 MyInt1 = MyInt1(t)
	var i2 MyInt2 = t

	fmt.Println(i1, i2)

	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 4)
	fmt.Printf("%v, %v", v, err)
}
