package main
/**
对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
对于一个协程调用，它的实参是在此协程被创建的时候估值的。
 */
import "fmt"

func main()  {
	func(){
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i) //先进后出
			//fmt.Println("a:", i)
		}
	}()

	fmt.Println()

	func(){
		for i := 0; i < 3; i++  {
			defer func() {
				fmt.Println("b:",i)
			}()
		}
	}()
}
