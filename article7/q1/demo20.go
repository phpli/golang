package main

import "fmt"

func main()  {

	ch1 := make(chan int,3)
	ch1 <- 2 //进入通道的不是２值本身，而是２的副本,放置值到通道，然后删除原值
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	fmt.Printf("The first element received from channel ch1:%v\n",elem1)
}
