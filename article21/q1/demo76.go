package main

import "fmt"

func main()  {
	m := map[string]int{"C": 1972, "C++": 1983, "Go": 2009}
	for lang, year := range m{
		fmt.Println(lang, year)
		fmt.Printf("%v : %v \n",lang, year)
	}

	a := [...]int{2,3,5,7,11}
	for i, prime := range a{
		fmt.Printf("%v :%v \n", i, prime)
	}

	s := []string{"go","defer","goto","var"}
	for i, keyword := range s {
		fmt.Printf("%v: %v \n", i, keyword)
	}
	aContainer := [5]int{1,2,3,4,5}
	// 忽略键值循环变量。
	for _, element := range aContainer {
		fmt.Printf("%v \n",element)
		// ...
	}

	// 忽略元素循环变量。
	for key, _ := range aContainer {
		element := aContainer[key]
		fmt.Printf("%v \n",element)
		// ...
	}

	// 舍弃元素循环变量。此形式和上一个变种等价。
	for key := range aContainer {
		element := aContainer[key]
		fmt.Printf("%v \n",element)
		// ...
	}

	// 键值和元素循环变量均被忽略。
	for _, _ = range aContainer {
		// 这个变种形式没有太大实用价值。
	}

	// 键值和元素循环变量均被舍弃。此形式和上一个变种等价。
	for range aContainer {
		// 这个变种形式没有太大实用价值。
	}
}
