package main

import "fmt"
//在下面的例子中，假设from（包括）和to（不包括）是两个合法的下标，并且from不大于to。
func main()  {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	s1 := s
	s2 := s
	// 第一种方法（保持剩余元素的次序）：
	s = append(s[:2],s[5:]...)
	fmt.Println(s)
	// 第二种方法（保持剩余元素的次序）：
	fmt.Println(s1)

	s1 = s1[:1 + copy(s1[5:], s1[8:])]
	fmt.Println(s1)

	// 第二种方法（保持剩余元素的次序）：

	if n := 8; len(s2)-2 < n {
		copy(s2[2:10], s2[10:])
	} else {
		copy(s2[2:10], s2[len(s2)-n:])
	}
	s2 = s2[:len(s2)-8]
	fmt.Println(s2)

	// 前弹出（pop front，又称shift）
	s, e := s[1:], s[0]
	// 后弹出（pop back）
	s, e = s[:len(s)-1], s[len(s)-1]
	// 前推（push front）
	s = append([]T{e}, s...)
	// 后推（push back）
	s = append(s, e)

}