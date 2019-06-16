package main

import (
	"errors"
	"fmt"
)

func main()  {
	fmt.Println("Enter function main.")
	caller()
	fmt.Println("Exit function main.")
}

func caller()  {
	fmt.Println("Enter function caller.")
	panic(errors.New("something wrong")) //正确例子
	panic(fmt.Println)//错误例子
	fmt.Println("Exit function caller.")
}