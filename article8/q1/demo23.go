package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	//示例１
	//只能发不能收的通道
	//如果这个操作符紧挨在chan的右边，那么就说明该通道只能发不接
	var uselessChan  = make(chan<- int ,1)

	//只能收不能发的通道
	var anotherUselessChan = make(<-chan int,1)
	//如果这个操作符紧挨在chan的左边，那么就说明该通道只能收不接

	fmt.Printf("The useless channels: %v,%v\n",uselessChan,anotherUselessChan)

	//示例２
	intChan1 := make(chan int,3)
	SendInt(intChan1)

	//示例４
	intChan2 := getIntChan()

	for elem := range intChan2{
		fmt.Printf("The element in intChan2: %v \n",elem)
	}

	//示例５
	_ = GetIntChan(getIntChan)

}

func SendInt(ch chan<- int)  {
	ch<- rand.Intn(1000)
}

//示例４
func getIntChan()<-chan int{
	num :=5
	ch := make(chan int,num)
	for i := 0; i<num ;i++  {
		ch<-i
	}
	close(ch)
	return ch
}

//示例５

type GetIntChan func()<-chan int