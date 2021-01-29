package main

import "fmt"

func main()  {
	const World  = "world"
	var hello  = "hello"
	//+ +=拼接
	var helloWorld = hello + "" + World
	helloWorld +=""
	fmt.Println(helloWorld)

	//字符串比较
	fmt.Println(hello == "world")
	fmt.Println(hello > helloWorld)
	fmt.Println(hello < helloWorld)
}
