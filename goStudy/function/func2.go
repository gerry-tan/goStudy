package main

func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func test2() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

func main() {
	//延迟求值
	funs := test()
	for _, f := range funs {
		f()
	}

	funs2 := test2()
	for _, f := range funs2 {
		f()
	}
}
