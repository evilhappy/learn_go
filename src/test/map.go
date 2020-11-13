/**
 * @Author: luhongguang
 * @Description:
 * @File:  test_map
 * @Date: 2020/10/20 下午4:48
 */
package main

import "fmt"

func main()  {
	oneClassStuTableNum := 25
	classStatic := [10]int{1, 2, 23, 10, 25, 35, 16, 37, 123, 235}
	classIdInfoList := make(map[int][]int, oneClassStuTableNum)
	for i := 0; i < 10; i++ {
		classId := classStatic[i]
		tableNum := classId % oneClassStuTableNum
		classIdInfoList[tableNum] = append(classIdInfoList[tableNum], classId)
	}
	fmt.Println(classIdInfoList)
}