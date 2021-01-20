package main

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x,y int
}

func main()  {
	for i :=0; i<=10; i++{
		fmt.Println(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Print(i)
		i := i // 这里声明的变量i遮挡了上面声明的i。
		// 右边的i为上面声明的循环变量i。
		i = 10 // 新声明的i被更改了。
		fmt.Print(i)
		//_ = i
	}

	if h := time.Now().Hour(); h<12{
		fmt.Println("现在为上午.")
	}else if h > 19{
		fmt.Println("现在为晚上")
	}else {
		fmt.Println("现在为下午")
		h := h
		_ = h
	}


	p := point{1,2}

	fmt.Printf("%v\n",p);
	fmt.Printf("%+v\n",p);
	fmt.Printf("%#v\n",p);

	rand.Seed(time.Now().UnixNano())
	//n := rand.Intn(3);
	//fmt.Println(n);
	switch n := rand.Intn(3); n {
	case 0: fmt.Println("n == 0")
	default: fmt.Println("n == 2")
	case 1: fmt.Println("n == 1")
	}

	for i := 90; i < 100; i++ {
		n := FindSmallestPrimeLargerThan(i)
		fmt.Print("最小的比", i, "大的素数为", n)
		fmt.Println()
	}
}

func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++{
		for i := 2; ; i++ {
			switch {
			case i * i > n:
				break Outer
			case n % i == 0:
				continue Outer
			}
		}
	}
	return n
}
