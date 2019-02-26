package main

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func read(path string) string {
	logs.Info("read startTime: " + strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	fi, err := os.Open(path)
	if err != nil {
		logs.Info(err)
	}
	defer fi.Close()
	fd, _ := ioutil.ReadAll(fi)
	logs.Info("read endTime: " + strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	return string(fd)
}

func write(path, data string) {
	logs.Info("write startTime: " + strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0766)
	if err != nil {
		logs.Error(err)
	}
	defer file.Close()
	file.Write([]byte(data))
	logs.Info("write endTime: " + strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
}

func main() {
	path := "/Users/tanqian/go/src/goStudy/files/test.txt"
	data := read(path)
	write("/Users/tanqian/go/src/goStudy/files/test2.txt", data)
}
