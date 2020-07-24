package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("字符串常用操作")

	fmt.Println("int 类型所占位数")
	fmt.Println(strconv.IntSize)

	fmt.Println("将字符串转换为 int 类型")
	var trastr01 string = "100"
	traint01, err_int := strconv.Atoi(trastr01)
	if err_int != nil {
		fmt.Println(err_int)
	} else {
		fmt.Println(traint01)
	}

	fmt.Println("将字符串转换为 float64 类型")
	var trastr02 string = "100.55"
	trafloat01, err_float := strconv.ParseFloat(trastr02, 10)
	if err_float != nil {
		fmt.Println(err_float)
	} else {
		fmt.Println(trafloat01)
	}

	fmt.Println("int 转化为字符串")
	trastr03 := strconv.Itoa(99)
	//trastr03 := string(100)
	fmt.Println("int 转化为字符串 " + trastr03)

	fmt.Println("字符串比较")
	var str01 string = "hello world"
	var str02 string = "hello world"
	//str02 := "世界你好"
	com01 := strings.Compare(str01, str02)
	if com01 == 0 {
		fmt.Println("相等")
	} else {
		fmt.Println("不相等 " + string(com01))
	}
	fmt.Println(com01)

	fmt.Println("包含")
	var isCon bool = strings.Contains(str01, "hello")
	fmt.Println(isCon)

	fmt.Println("查找位置")
	var theIndex int = strings.Index(str01, "w")
	fmt.Println(theIndex)
	fmt.Println(strings.Index(str01, "w"))
	fmt.Println(str01[0:theIndex])
	lastIndex := strings.LastIndex(str01, "o")
	fmt.Println("此字段在字符串中最后出现位置的索引 " + strconv.Itoa(lastIndex))

	fmt.Println("统计给定子串sep的出现次数，sep为空时候返回1 + 字符串长度")
	fmt.Println(strings.Count("cheeseeee", "ee")) //3
	fmt.Println(strings.Count("cheese", ""))      //7

	fmt.Println("重复字符串指定次数，最后返回新生成的重复的字符串")
	fmt.Println("hello" + strings.Repeat("world", 2))

	fmt.Println("替换")
	var str03 string = "/Users/hello/works/admin/"
	str04 := strings.Replace(str03, "/", "***", -1)
	str05 := strings.Replace(str03, "/", "***", 3)
	fmt.Println(str04)
	fmt.Println(str05)

	fmt.Println("删除字符串的开头和结尾")
	fmt.Println("删除两头的 / = " + strings.Trim(str03, "/"))
	fmt.Println("删除左边的 / = " + strings.TrimLeft(str03, "/"))
	fmt.Println("删除右边的 / = " + strings.TrimRight(str03, "/"))

	fmt.Println("删除收尾空格")
	str06 := " hello good day "
	fmt.Println(strings.TrimSpace(str06))

	fmt.Println("大小写")
	str07 := "hi girl"
	fmt.Println(strings.Title(str07))
	fmt.Println(strings.ToLower(str07))
	fmt.Println(strings.ToUpper(str07))

	fmt.Println("是否有 前缀/后缀")
	fmt.Println(strings.HasPrefix("fix_bug", "fix"))
	fmt.Println(strings.HasSuffix("Golang", "ang"))

	fmt.Println("字符串分割")
	fieldsStr := "  hello   it's  a  nice day today    "
	fieldSlice := strings.Fields(fieldsStr)
	fmt.Println(fieldSlice)
	for i, v := range fieldSlice {
		fmt.Printf("下标 %d 对应值 = %s \n", i, v)
	}
	for i := 0; i < len(fieldSlice); i++ {
		fmt.Println(fieldSlice[i])
	}

	fmt.Println("讲元素类型是string的slice，使用分隔符号来拼接组成一个字符串，比如逗号分割")
	str09 := strings.Join(fieldSlice, ",")
	fmt.Println(str09)


	fmt.Println("逗号分割的字符串转为切片")
	str08 := "a,b,c,d,e,f"
	sliceList := strings.Split(str08, ",")
	fmt.Println(sliceList)
	for k, v := range sliceList {
		fmt.Println(k, v)
	}


}
