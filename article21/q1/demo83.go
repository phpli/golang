package main

import (
	"fmt"
	"strings"
)

func main()  {
	var helloWorld = "hello world"
	var hello = helloWorld[:5]
	fmt.Println(hello[0])
	fmt.Println("%T \n", hello[0])
	// hello[0]是不可寻址和不可修改的，所以下面
	// 两行编译不通过。
	/*
		hello[0] = 'H'         // error
		fmt.Println(&hello[0]) // error
	*/

	// 下一条语句将打印出：5 12 true

	fmt.Println(len(hello), len(helloWorld), strings.HasPrefix(helloWorld, hello))
}
