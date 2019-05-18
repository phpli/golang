package main

import "fmt"

var container = []string{"zero","one","two"}

func main()  {
	container := map[int]string{0:"zero",1:"one",2:"two"}

	//方式１
	_,ok1 := interface{}(container).([]string)
	_,ok2 := interface{}(container).(map[int]string)

	if!(ok1 || ok2) {
		fmt.Printf("Error:unsupported containter type:%T\n",container)
		return
	}

	//方式２
	elem,err := getElement(container)

	if err!=nil{
		fmt.Printf("Error:%s!\n",err)
	}
	fmt.Printf("The element is %q.(container type :%T)\n",elem,container)

}

func getElement(containerI interface{}) (elem string,err error) {
	switch t := containerI.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported containter type :%T",containerI)
		return
	}
	return
}
