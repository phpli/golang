package main

import "fmt"

func main()  {
	defer fmt.Println("first defer")
	for i := 0;i < 3 ;i++  {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}
//defter 先进后出（FILO）
//结果
//last defer
//defer in for [2]
//defer in for [1]
//defer in for [0]
//first defer
