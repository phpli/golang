package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var x struct{
		a int64
		b bool
		c string
	}

	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)

	fmt.Println(M, N)
	fmt.Println(unsafe.Alignof(x.a))
	fmt.Println(unsafe.Alignof(x.b))
	fmt.Println(unsafe.Alignof(x.c))

	fmt.Println(unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Offsetof(x.c))

	type T struct {
		d string
	}

	type S struct {
		e bool
	}

	var y struct{
		xx int64
		*S
		T
	}
	//reflect.SliceHeader{}
	fmt.Println(unsafe.Offsetof(y.xx))
	fmt.Println(unsafe.Offsetof(y.S))
	fmt.Println(unsafe.Offsetof(y.T))
	fmt.Println(unsafe.Offsetof(y.d))
	//fmt.Println(unsafe.Offsetof(y.e)) // error
	fmt.Println(unsafe.Offsetof(y.S.e))
}