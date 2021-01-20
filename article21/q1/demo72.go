package main



import (
	"fmt"
	"time"
)

/**
对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
对于一个协程调用，它的实参是在此协程被创建的时候估值的。
 */

func main()  {
	var a = 123
	b := 1
	go func(x int,b int) {
		time.Sleep(time.Second)
		fmt.Println(x, a, b)
	}(a,b)

	a = 789
	time.Sleep(2*time.Second)
}
