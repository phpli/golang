package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
	}
	persons := [2]Person {{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons{
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[1].name = "Jack"
		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", &persons)

	personss := []Person {{"Alice", 28}, {"Bob", 25}}//切片改变的是底层共享数据
	for i, p := range personss {
		fmt.Println(i, p) //但是此时打印已经发生变化
		// 这次，此修改将反映在此次遍历过程中。
		personss[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。 git11
		p.age = 31
	}
	fmt.Println("persons:", &personss)

	var a  [100]int
	for i,n := range &a{ // 复制一个指针的开销很小
		fmt.Println(i,n)
	}

	for i, n:= range a[:]{ // 复制一个切片的开销很小
		fmt.Println(i, n)
	}

	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	for i, n := range p { // panic
		fmt.Println(i, n)
	}
}