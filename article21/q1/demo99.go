package main

import (
	"fmt"
	"unsafe"
)

//使用模式三：将一个非类型安全指针转换为一个uintptr值，然后此uintptr值参与各种算术运算，再将算术运算的结果uintptr值转回非类型安全指针。

type To struct {
	x bool
	y [3]int16
}

const N  =  unsafe.Offsetof(To{}.y)
const M  =  unsafe.Offsetof(To{}.y[0])

func main()  {
	t := To{
		y: [3]int16{123, 456, 789},
	}

	p := unsafe.Pointer(&t)
	//分开写可能导致 变量被回收
	ty2 := (*int16)(unsafe.Pointer(uintptr(p)+N+M+M))
	fmt.Println(*ty2) // 789
}