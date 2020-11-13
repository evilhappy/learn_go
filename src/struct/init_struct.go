/**
 * @Author: luhongguang
 * @Description:
 * @File:  init_struct
 * @Date: 2020/11/13 上午11:29
 */
package main

import "fmt"

type Cat struct {
	Name string `json:"name"`
	High string `json:"high"`
	Sex  string `json:"sex"`
}

func main()  {
	cat := new(Cat)
	cat.Name = "hello"
	cat.High = "12"

	fmt.Printf("%+v",cat)
}