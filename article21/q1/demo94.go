package main

import "fmt"

//一个[]T类型的值不能直接被转换为类型[]I，即使类型T实现了接口类型I
//比如，我们不能直接将一个[]string值转换为类型[]interface{}。 我们必须使用一个循环来实现此转换：

type I interface {
	m(int)bool
}

type T string

func (t T) m(n int) bool  {
	return len(t) > n
}

func main()  {
	words := []string{
		"Go", "is", "a", "high",
		"efficient", "language.",
	}

	iw := make([]interface{}, 0, len(words))

	for _, w := range words{
		iw = append(iw, w)
	}
	fmt.Println(iw...)

	var i I = T("gopher")
	fmt.Println(i.m(5))                          // true
	fmt.Println(I.m(i, 5))                       // true
	fmt.Println(interface {m(int) bool}.m(i, 5)) // true

	// 下面这几行被执行的时候都将会产生一个恐慌。
	I(nil).m(5)
	I.m(nil, 5)
	interface {m(int) bool}.m(nil, 5)

}
