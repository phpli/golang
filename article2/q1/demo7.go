package main

import (
	"flag"
	"fmt"
)

func main()  {
	var name  string

	flag.StringVar(&name,"name","everyone","The greeting object!")

	//方式１
	var name2 = flag.String("name2","everyone","The greeting object!")

	//方式２
	name1 := flag.String("name1","everyone","The greeting object!")

	flag.Parse()

	fmt.Printf("Hello,%s!\n%s!\n%s!\n",name,*name1,*name2)

	//fmt.Printf("Hello,%s!\n",*name1)

	//fmt.Printf("Hello,%s!\n",*name2)

}