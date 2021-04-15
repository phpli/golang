package main

import (
	"fmt"
	"reflect"
	"unsafe"
)
//一个使用了reflect.SliceHeader的例子：
func main()  {
	a := [6]byte{'G', 'o', '1', '0', '1'}
	bs := []byte("Golang")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hdr.Data = uintptr(unsafe.Pointer(&a))

	hdr.Len = 2
	hdr.Cap = len(a)
	fmt.Printf("%s\n", bs)
	bs = bs[:cap(bs)]
	fmt.Printf("%s\n", bs)

	//一般说来，我们只应该从一个已经存在的字符串值得到一个*reflect.StringHeader指针，
	// 或者从一个已经存在的切片值得到一个*reflect.SliceHeader指针， 而不应该从一个StringHeader值生成一个字符串，或者从一个SliceHeader值生成一个切片。 比如，下面的代码是不安全的：
	//var hdr reflect.StringHeader
	//hdr.Data = uintptr(unsafe.Pointer(new([5]byte)))
	//// 在此时刻，上一行代码中刚开辟的数组内存块已经不再被任何值
	//// 所引用，所以它可以被回收了。
	//hdr.Len = 5
	//s := *(*string)(unsafe.Pointer(&hdr)) // 危险！
}