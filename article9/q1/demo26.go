package main

import "fmt"

//都是函数声明，其实是一个函数
type Printer func(contents string) (n int, err error)

func printToStd(contents string)(bytesNum int,err error)  {
	return  fmt.Println(contents)
}

func main()  {
	var p Printer
	p = printToStd
	p("something")
}
