package main

import (
	"flag"
	"fmt"
)

//var name = flag.String("name", "everyone", "a greeting object")
var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "a greeting object")
}

func main() {
	// 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
