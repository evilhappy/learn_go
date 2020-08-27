package main

import (
	"flag"
	"fmt"
	"os"
)

//var name = flag.String("name", "everyone", "a greeting object")
var name1 string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name1, "name", "everyone", "a greeting object")
}

func main() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	// 函数flag.Parse用于真正解析命令参数，并把它们的值赋给相应的变量
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name1)
}
