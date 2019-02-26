package main

import (
	"fmt"
	"github.com/astaxie/beego/utils"
	"github.com/magiconair/properties"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func getAppPath() string {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return appPath
}

func test() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {

	appPath := getAppPath()
	fmt.Println("appPath: ", appPath)

	workPath, err := os.Getwd()
	fmt.Println("workPath: ", workPath)

	appConfigPath := filepath.Join(appPath, "props", "db.properties")

	exists := utils.FileExists(appConfigPath)
	fmt.Println(exists)

	props, err := properties.LoadFile("goStudy/props/db.properties", properties.UTF8)
	if err != nil {
		panic(err)
	}

	value, ok := props.Get("db.mysql.password")
	if !ok {
		return
	}

	fmt.Println(value)
}
