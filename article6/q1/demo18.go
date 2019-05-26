package main

func main()  {
	//示例１
	//var badMap1 = map[[]int]int{} //// 这里会引发编译错误。
	//_ = badMap1
	//
	////示例２
	//var badMap2 = map[interface{}]int{
	//	"1" :1,
	//	[]int{2}:2,
	//	3:3,
	//}
	//_ = badMap2
	//
	////示例３
	//var badMap3  map[[1][]string]int // 这里会引发编译错误。
	//_ = badMap3
	//
	//
	////示例４
	//var badMap５  map[[1][2][3][]string]int //这里会报错
	//
	//_ = badMap５
	//
	////示例５
	//type  BadKey1 struct {
	//	slice []string
	//}
	//var badMap4 map[BadKey1]int // 这里会引发编译错误。
	//
	//_ = badMap4
	//
	////示例６
	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//
	//var badMap6  map[BadKey2]int // 这里会引发编译错误。
	//
	//_ = badMap6

}
