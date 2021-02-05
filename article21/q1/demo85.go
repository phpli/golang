package main

import "fmt"

func Sum(values ...int64)(sum int64){
	sum = 0
	for _, v := range values{
		sum += v
	}
	return
}

func main()  {
	a0 := Sum()
	a1 := Sum(2)
	a2 := Sum(2, 3, 4)

	b0 := Sum([]int64{}...)
	b1 := Sum([]int64{2}...)
	b2 := Sum([]int64{2, 3, 4}...)
	fmt.Println(a0, a1, a2) // 0 2 10
	fmt.Println(b0, b1, b2) // 0 2 10

	tokens := []string{"Go", "C", "Rust"}
	langsA := Concat(",", tokens...)
	langsB := Concat(",", "Go", "C", "Rust")
	fmt.Println(langsA == langsB)
}
//所有的函数调用的传参均属于值复制
func Concat(sep string, tokens ...string)(r string){
	for i, t := range tokens {
		if i !=0 {
			r += sep
		}
		r += t
	}
	return
}