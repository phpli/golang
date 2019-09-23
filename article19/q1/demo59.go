package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"sync"
	"time"
)

// singleHandler 代表单次处理函数的类型。
type singleHandler func()(data string, n int, err error)

//handlerConfig 代表处理的配置的流程
type handlerConfig struct {
	handler 	singleHandler	//单次处理函数
	goNum 		int			  	// 需要启用goroutine的数量
	number  	int			  	// 单个goroutine的处理次数
	interval	time.Duration	//单个goroutine 的执行时间
	counter		int				// 数据量计数器，以字节为单位。
	counterMu	sync.Mutex		// 数据量计数器专用的互斥锁。
}

//count 会增加计数器的值，并返回增加后的计数
func (hc *handlerConfig) count(increment int) int {
		hc.counterMu.Lock()
		defer hc.counterMu.Unlock()
		hc.counter += increment
		return hc.counter
}

func main()  {
	// mu 代表以下流程要使用的互斥锁。
	// 在下面的函数中直接使用即可，不要传递。
	var mu sync.Mutex
	// genWriter 代表的是用于生成写入函数的函数。
	genWriter := func(writer io.Writer) singleHandler{
		return func() (data string, n int, err error) {
				//准备数据
				data = fmt.Sprintf("%s\t",
					time.Now().Format(time.StampNano))
				//写入数据
				mu.Lock()
				defer mu.Unlock()
				n, err = writer.Write([]byte(data))
				n = len(data)
				return
		}
	}

	genReader := func(reader io.Reader) singleHandler{
		return func() (data string, n int, err error) {
				buffer, ok := reader.(*bytes.Buffer)
				if !ok{
					err = errors.New("unsupported reader")
					return
				}
				//读取数据
				mu.Lock()
				defer mu.Unlock()
			data, err = buffer.ReadString('\t')
			n = len(data)
			return
		}
	}
	//buffer 代表缓冲区。
	var buffer bytes.Buffer
	// 数据写入配置。
	writingConfig := handlerConfig{
		handler:   genWriter(&buffer),
		goNum:     5,
		number:    4,
		interval:  time.Microsecond*100,
	}

	// 数据读取配置。
	readingConfig := handlerConfig{
		handler:   genReader(&buffer),
		goNum:     10,
		number:    2,
		interval:  time.Microsecond*100,
	}

	// sign 代表信号的通道。
	sign := make(chan struct{}, writingConfig.goNum + readingConfig.goNum)

	for i := 1; i < writingConfig.goNum ; i++  {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1 ; j < writingConfig.number ; j++  {
					time.Sleep(writingConfig.interval)
					data , n, err := writingConfig.handler()
					if err != nil{
							log.Printf("writer [%d-%d]: error: %s",
								i,j,err)
							continue
					}
					total := writingConfig.count(n)
					log.Printf("writer [%d-%d]: %s (total: %d)",
						i, j, data, total)
			}
		}(i)
	}

	// 启用多个goroutine对缓冲区进行多次数据读取。
	for i := 1; i <= readingConfig.goNum ; i++  {
		go func(i int) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= readingConfig.number ; j++  {
					time.Sleep(readingConfig.interval)
					var data string
					var n int
					var err error
					for {
							data , n, err = readingConfig.handler()
							if err == nil || err != io.EOF{
								break
							}
						// 如果读比写快（读时会发生EOF错误），那就等一会儿再读。
						time.Sleep(readingConfig.interval)
					}
					if err != nil{
						log.Printf("reader [%d-%d]: error: %s",
							i, j, err)
						continue
					}
					total := readingConfig.count(n)
					log.Printf("reader [%d-%d]: %s (total: %d)",
						i, j, data, total)
			}
		}(i)
	}

	//
	signNumber := writingConfig.goNum + readingConfig.goNum
	for j := 0; j < signNumber; j++ {
		<-sign
	}
}


