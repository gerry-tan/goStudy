package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

type MyStruct struct {
	Name string
	Age  int8
}

func main() {
	// 创建一个默认过期时间为5分钟，清理间隔时间为10分钟的高速缓存
	c := cache.New(5*time.Second, 10*time.Second)

	// 设置“name”键的值为“Jame”，默认过期时间
	c.Set("name", "Jame", cache.DefaultExpiration)

	// 设置“baz”为42，不过期
	// 如果没有重置或者删除的话，它不会被删除
	c.Set("baz", 42, cache.NoExpiration)

	// 获取"foo"对应的字符串
	name, found := c.Get("name")
	if found {
		fmt.Println(name)
	}

	// 因为Go是一种静态类型语言，而cache可以存储任何类型，因此可以使用断言来判断任意类型
	/*foo, found := c.Get("foo")
	if found {
		MyFunction(foo.(string))
	}*/

	obj := &MyStruct{
		Name: "Jame",
		Age:  18,
	}
	fmt.Println("obj: ", obj)
	// 需要高性能？那就存指针吧
	c.Set("foo", obj, cache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo := x.(*MyStruct)
		fmt.Println("value: ", x, foo)
	}

}
