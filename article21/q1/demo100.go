package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
//一个使用reflect.StringHeader的例子：
func main()  {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g'}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s)
	// 现在，字符串s和切片a共享着底层的byte字节序列，
	// 从而使得此字符串中的字节变得可以修改。
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	fmt.Println(s) // Google
}
