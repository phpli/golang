package main

import (
	"flag"
	"github.com/phpli/golang/article1/q2/lib"
)

var name  string

func init()  {
	flag.StringVar(&name,"name","everyone","The greeting object.")
}

func main()  {
	flag.Parse()
	//beego.Run()
	lib.Hello(name)
}
