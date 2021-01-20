package main

import (
	"sync"
	"time"
)

var wgs sync.WaitGroup

func main()  {
	wgs.Add(1)
	go func() {
		time.Sleep(time.Second*2)
		wgs.Wait()
	}()
	wgs.Wait()
}
