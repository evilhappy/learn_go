package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 3, 5, 7, 9}
	reverse(a[:]) //注意这里参数
	fmt.Println(a)

	reverse(a[2:])
	fmt.Println(a)

	reverse(a[:2])
	fmt.Println(a)
}
