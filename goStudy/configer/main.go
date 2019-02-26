package main

import (
	"fmt"
	"github.com/itanqian/configor"
	"os"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
}{}

func main() {
	filepath, _ := os.Getwd()
	fmt.Println(filepath)

	configor.Load(&Config, "/Users/tanqian/go/src/goStudy/configer/conf.yml")

	fmt.Printf("%v\n", Config)
}
