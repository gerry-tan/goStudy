package main

import (
	"fmt"
	"reflect"
)

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	//结构体属性相同,顺序相同则相等
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//结构体中含有 map,slice 不能直接比较
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}

}
