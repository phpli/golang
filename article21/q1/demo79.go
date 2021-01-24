package main

import "fmt"

func main()  {
	//我们可以通过数组的指针来访问和修改此数组中的元素。如果此指针是一个nil指针，将导致一个恐慌。
	a := [5]int{2,3,5,7,11}
	p := &a
	p[0], p[1] = 17,19
	fmt.Println(a)
	//我们可以从一个数组的指针派生出一个切片。从一个nil数组指针派生切片将导致一个恐慌。
	b := [6]int{1,2,3,4,5,6}
	s := b[1:3]
	fmt.Println(s)
	//b = nil
	//s = b[0:0]

	//内置len和cap函数调用接受数组指针做为实参。 nil数组指针实参不会导致恐慌。
	var pa *[5]int
	fmt.Println(len(pa),cap(pa))

}