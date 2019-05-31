package main

import "fmt"

var channels = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var numbers  = []int{1,2,3}

func main()  {
	select {
	case getChan(0) <- getNumber(0):
		fmt.Printf("The first candidate is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Printf("The second candidate is selected.")
	case getChan(2) <- getNumber(2):
		fmt.Printf("The third candidate is selected.")
	default:
		fmt.Printf("No candidate is selected.")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n",i)
	return numbers[i]
}

func getChan(i int) chan int{
	fmt.Printf("channels[%d]\n",i)
	return channels[i]
}


