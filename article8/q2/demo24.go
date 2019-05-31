package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//示例１
	example1()
	example2()
}


//示例１
func example1()  {
	//准备好几个通道
	intChannels := [3]chan int{
			make(chan int, 1),
			make(chan int, 1),
			make(chan int, 1),
	}
	//随机选择一个并向他发送元素值
	index := rand.Intn(3)
	fmt.Printf("The index :%d\n",index)
	intChannels[index] <- index
	//哪一个通道中元素，哪一个ｃａｓｅ就会被执行
	select {
	case <-intChannels[0]:
		fmt.Printf("The first candidate case is selected.")
	case <-intChannels[1]:
		fmt.Printf("The second candidate case is selected.")
	case <-intChannels[2]:
		fmt.Printf("The third candidate case is selected.")
	default:
		fmt.Printf("No candidate case is selected.")
	}
}


//示例２
func example2()  {
	intChan := make(chan int, 1)
	//一秒后关闭通道
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _,ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}
