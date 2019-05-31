package main

import (
	"errors"
	"fmt"
)

type operate func(x,y int)int

//方案１

//原数组不会改变。为什么呢？原因是，所有传给函数的参数值都会被函数在其内部使用的并不是参数值的原值，而是它的副本。
//

func calculate(x int, y int, op operate)(int,error)  {
	if op==nil {
			return 0, errors.New("invalid operation")
	}
	return op(x,y),nil
}


//方案２

type calculateFunc func(x int, y int)(int,error)

func getCalculator(op operate) calculateFunc{
	return func(x int, y int) (i int, e error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x,y),nil
	}
}

func main()  {
	//方案１
	x,y := 12, 23
	op := func(x,y int)int {
		return x + y
	}
	result,err := calculate(x,y,op)
	fmt.Printf("The result :%d(error:%v)\n", result,err)

	result, err = calculate(x,y,nil)
	fmt.Printf("The result :%d(error:%v)\n", result,err)

	//方案２
	x,y = 56, 78

	add := getCalculator(op)
	result,err = add(x, y)
	fmt.Printf("The result :%d(error:%v)\n", result, err)
}
