package main

import "fmt"

func main()  {
		//示例１
		numbers1 := []int{1,2,3,4,5,6}
		for i := range numbers1{
			if 3 ==i{
					numbers1[i] |= i
			}
		}
		fmt.Println(numbers1)
		fmt.Println()

		//示例２
		numbers2 := [...]int{1,2,3,4,5,6}
		maxInde2 := len(numbers2)-1
		for i,e := range numbers2{  // i 为索引，e 为该索引的元素值
			if i == maxInde2{
				numbers2[0] += e
			}else {
				numbers2[i+1] += e
			}
		}

		fmt.Println(numbers2)
		fmt.Println()

		//示例３
		numbers3 := []int{1,2,3,4,5,6}
		maxInde3 := len(numbers2)-1
		for i,e := range numbers3{
			if i == maxInde3{
				numbers3[0] += e
			}else {
				numbers3[i+1] += e
			}
		}
		fmt.Println(numbers3)
}
