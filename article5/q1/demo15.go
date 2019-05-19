package main

import "fmt"

func main()  {
	//示例１
	s1 := make([]int,5)
	fmt.Printf("The length of s1:%d\n",len(s1))
	fmt.Printf("The capacity of s1:%d\n",cap(s1))
	fmt.Printf("The value of s1:%d\n",s1)
	s2 := make([]int,5,8)
	fmt.Printf("The length of s2:%d\n",len(s2))
	fmt.Printf("The capacity of s2:%d\n",cap(s2))
	fmt.Printf("The value of s2:%d\n",s2)
	fmt.Println()

	//示例２
	s3 := []int{1,2,3,4,5,6,7,8}
	s4 := s3[3:6]//从索引三开始到索引六，【ａ,b】从索引ａ到索引b，而且还不包括结束索引Ｂ
	fmt.Printf("The length of s4:%d\n",len(s4))
	fmt.Printf("The capacity of s4:%d\n",cap(s4))//容量等于cap(s3)-3 底层数组的容量减去切片的起始索引
	fmt.Printf("The value fo s4:%d\n",s4)
	fmt.Println()

	//示例３
	s5 := s4[:cap(s4)] //s4[0:cap(s4)]顺便提一下把切片的窗口向右扩展到最大的方法
	fmt.Printf("The length of s5:%d\n",len(s5))
	fmt.Printf("The capacity of s5:%d\n",cap(s5))
	fmt.Printf("The value fo s5:%d\n",s5)
}
