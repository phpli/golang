package main

import "fmt"

func main()  {
	type Ta []int
	type Tb []int
	des := Ta{1,2,3}
	src := Tb{5,6,7,8,9}
	n := copy(des, src) // copy 返回的数据 是复制元素数量  copy函数的第一个参数为目标切片，第二个参数为源切片。
	//此时 des{5,6,7}

	fmt.Println(n,des,cap(des))
	fmt.Println(n,src)

	fmt.Println(des[1:]) //des[1:] -->6 7
	fmt.Println(des)//此时 des{5,6,7}
	n = copy(des[1:], des[0:]) // des 567 去覆盖->67 此时修改的是 des 的底层des元数据
	fmt.Println(des)//为啥变成[5,5,6] --->
	fmt.Println(des[1:],cap(des[1:]))

	a :=[4]int{}
	n = copy(a[:], src)
	fmt.Println(n, a)
	n = copy(a[:], a[2:])
	fmt.Println(n,a)
}
