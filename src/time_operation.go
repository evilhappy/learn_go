package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("时间")
	now := time.Now()
	//这个 2006-01-02 15:04:05 数字是一定的，如果是12小时制的话，15改为03
	fmt.Println("当前时间")
	formatStart := now.Format("2006-01-02 15:04:05")
	fmt.Println(formatStart)
	fmt.Println("时间戳")
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)


	fmt.Println("一分钟前")
	m,_ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1.Format("2006-01-02 15:04:05"))

	fmt.Println("两个小时前")
	h,_ := time.ParseDuration("-2h")
	h1 := now.Add(h)
	fmt.Println(h1.Format("2006-01-02 15:04:05"))

	fmt.Println("三天前")
	d1 := now.AddDate(0, 0, -3)
	fmt.Println(d1.Format("2006-01-02 15:04:05"))

	fmt.Println("十分钟后")
	mm,_ := time.ParseDuration("10m")
	mm1 := now.Add(mm)
	fmt.Println(mm1.Format("2006-01-02 15:04:05"))

	fmt.Println("两小时后")
	hh,_ := time.ParseDuration("2h")
	hh1 := now.Add(hh)
	fmt.Println(hh1.Format("2006-01-02 15:04:05"))

	fmt.Println("三天后")
	dd1 := now.AddDate(0, 0, 3)
	fmt.Println(dd1.Format("2006-01-02 15:04:05"))
}
