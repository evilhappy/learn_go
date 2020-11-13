/**
 * @Author: luhongguang
 * @Description:
 * @File:  get_params
 * @Date: 2020/11/12 下午7:36
 */
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

var help = func () {
	fmt.Println("Usage for calc tool.")
	fmt.Println("====================================================")
	fmt.Println("add 1 2, return 3")
	fmt.Println("sub 1 2, return -1")
	fmt.Println("mul 1 2, return 2")
	fmt.Println("sqrt 2, return 1.4142135623730951")
}


func CalcByOs() error {
	args := os.Args
	if len(args) < 3 || args == nil {
		help()
		return nil
	}
	operate := args[1]
	switch operate {
	case "add":{
		rt := 0
		number_one, err1 := strconv.Atoi(args[2])
		number_two, err2 := strconv.Atoi(args[3])
		if err1 == nil && err2 == nil {
			rt = number_one + number_two
			fmt.Println("Result ", rt)
		}
	}
	case "sub":
		{
			rt := 0
			number_one, err1 := strconv.Atoi(args[2])
			number_two, err2 := strconv.Atoi(args[3])
			if err1 == nil && err2 == nil {
				rt += number_one - number_two
				fmt.Println("Result ", rt)
			}
		}
	case "mul":
		{
			rt := 1
			number_one, err1 := strconv.Atoi(args[2])
			number_two, err2 := strconv.Atoi(args[3])
			if err1 == nil && err2 == nil {
				rt = number_one * number_two
				fmt.Println("Result ", rt)
			}
		}
	case "sqrt":
		{
			rt := float64(0)
			if len(args) != 3 {
				fmt.Println("Usage: sqrt 2, return 1.4142135623730951")
				return nil
			}
			number_one, err := strconv.ParseFloat(args[2], 64)
			if err == nil {
				rt = math.Sqrt(number_one)
				fmt.Println("Result ", rt)
			}
		}
	default:
		help()

	}
	return nil
}

func main() {
	err := CalcByOs()
	fmt.Println(err)
}