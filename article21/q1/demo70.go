package main

import "fmt"

func Triple(n int) (r int)  {
	defer func() {
		r += n
	}()
	r = n+n
	return
}

func main()  {
	fmt.Println(Triple(5))
}
