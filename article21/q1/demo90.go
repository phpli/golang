package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter用来删除重复的数字。
type UniqueFilter struct {}

func (UniqueFilter) About() string  {
	return  "删除重复的数字"
}

func (UniqueFilter) Process(inputs []int) []int  {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _,n := range inputs{
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

type MultipleFilter int

func (mf MultipleFilter) About() string  {
	return fmt.Sprintf("保留%v的倍数", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int  {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs  {
		if n % int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

func filteAndPrint(fltr Filter, unfiltered []int) []int {
	// 在fltr参数上调用方法其实是调用fltr的动态值的方法。
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About() + ":\n\t", filtered)
	return filtered
}

func main()  {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("过滤之前:\n\t", numbers)

	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	for _, fltr  := range filters {
		numbers = filteAndPrint(fltr, numbers)
	}
}

