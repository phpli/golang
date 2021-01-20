package main

import "fmt"

func main()  {
	defer fmt.Println("the third line")
	defer fmt.Println("the second line") //延迟堆栈 ，先进的后出
	fmt.Println("the first line")
}
