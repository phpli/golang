package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	type MyString string
	ms := []MyString{"c", "c++", "go"}
	fmt.Printf("%s\n", ms)

	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Rust"
	fmt.Printf("%s\n", ms)
	ms = *(*[]MyString)(unsafe.Pointer(&ss))
}
