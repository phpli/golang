package main

import "fmt"

func main()  {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1,2,3}, map[int]bool{},nil,
	}

	for _, x := range values{
		switch v := x.(type) {
		case []int: // 一个类型字面表示
			// 在此分支中，v的类型为string。
			fmt.Println("int slice:", v)
		case string://// 一个类型名
			fmt.Println("string:",v)
		case int, float64, int32:// 多个类型名
			fmt.Println("number:",v)
		case nil:
			fmt.Println(v)
		default:
			fmt.Println(v)
		}
		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
		//上面的内容等价于下面
	}
	fmt.Println("otherssss")
	for _, x := range values{
		if v, ok := x.([]int); ok{
			fmt.Println("int slice:", v)
		}else if v, ok := x.(string); ok{
			fmt.Println("string:", v)
		}else if x == nil{
			v := x
			fmt.Println(v)
		}else {
			_, isInt := x.(int)
			_, isFloat64 := x.(float64)
			_, isInt32 := x.(int32)
			if isInt || isFloat64 || isInt32{
				v := x
				fmt.Println(v)
			}else {
				v := x
				fmt.Println("others:", v)
			}
		}
	}
}