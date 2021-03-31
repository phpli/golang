package main

import "fmt"

type Writer interface {
	Writer(buf []byte) (int, error)
}

type DummyWriter struct {

}

func (DummyWriter) Writer(buf []byte) (int, error)  {
	return len(buf),nil
}

func main()  {
	var x interface{} = DummyWriter{}
	var y interface{} = "abc"
	var w Writer
	var ok bool

	// DummyWriter既实现了Writer，也实现了interface{}。
	w, ok = x.(Writer)
	fmt.Println(w, ok)
	x, ok = w.(interface{})
	fmt.Println(x, ok)

	// y的动态类型为string。string类型并没有实现Writer。
	w, ok = y.(Writer)
	fmt.Println(w, ok) // <nil> false
	//w = y.(Writer)     // 将产生一个恐慌
}
