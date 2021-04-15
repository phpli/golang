package main

import (
	"fmt"
	"unsafe"
)

//使用模式一：将类型*T1的一个值转换为非类型安全指针值，然后将此非类型安全指针值转换为类型*T2。
func main()  {
	type MyString string
	ms := []MyString{"c", "c++", "go"}
	fmt.Printf("%s\n", ms)

	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Rust"
	fmt.Printf("%s\n", ms)
	ms = *(*[]MyString)(unsafe.Pointer(&ss))
}
