package main

import (
	"flag"
	"golang/article1/q2/lib"
	"github.com/astaxie/beego"
)

var name  string

func init()  {
	flag.StringVar(&name,"name","everyone","The greeting object.")
}

func main()  {
	flag.Parse()
	beego.Run()
	lib.Hello(name)
}