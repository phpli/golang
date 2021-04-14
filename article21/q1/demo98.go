package main

import (
	"fmt"
	"unsafe"
)
//使用模式二：将一个非类型安全指针值转换为一个uintptr值，然后使用此uintptr值。
//此模式不是很有用。一般我们将最终的转换结果uintptr值输出到日志中用来调试，但是有很多其它安全并且简洁的途径也可以实现此目的。
func main()  {
	type T struct {
		a int
	}
	var t T
	fmt.Printf("%p\n", &t)
	println(&t)
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t)))
}
