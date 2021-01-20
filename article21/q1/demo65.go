package main

import (
	"log"
	"math/rand"
	"time"
)

func main()  {
	//log.Println(runtime.NumCPU());
	//return
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!",10)
	go SayGreetings("hello!",10)
	time.Sleep(2 * time.Second)
}

func SayGreetings(greetings string, times int)  {
	for i :=0; i < times; i++  {
		log.Println(greetings)
		d := time.Second*time.Duration(rand.Intn(5))/2
		time.Sleep(d)
	}
}
