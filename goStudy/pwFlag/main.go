package main

import (
	"flag"
	"fmt"
)

var (
	length  int
	charset string
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成秘钥的长度")
	flag.StringVar(&charset, "t", "num",
		"-t 指定密码生成的字符集, num: 纯数字; char: 纯英文字母; mix: 使用数字和字母; advance: 使用数字,字母以及特殊字符")
	flag.Parse()
}

func main() {
	parseArgs()
	fmt.Printf("length: %d, charset: %s", length, charset)
}
