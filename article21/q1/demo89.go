package main

import "fmt"

func main()  {
	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123
	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok)
	n = x.(int)
	fmt.Println(n)

	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok)
	// 情形三：
	a = x.(float64) // 将产生一个恐慌
}