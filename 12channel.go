package main

import (
	"log"
	"sync"
	"time"
)

// 声明并且初始化一个长度为0的int类型的channel
var myChan = make(chan int)
var myChan2 = make(chan string)
var doneChane = make(chan struct{})

func setTimeOut(name string, wait *sync.WaitGroup) {
	// log.Println("timeout-start,", name)
	time.Sleep(2 * time.Second)
	// log.Println("timeout-end,", name)

	myChan <- 1
	myChan2 <- name

	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	// log.Println("main function start")
	var startTime time.Time = time.Now()
	wait.Add(3)
	go setTimeOut("goroutine 1", &wait)
	go setTimeOut("goroutine 2", &wait)
	go setTimeOut("goroutine 3", &wait)

	go func() {
		defer func() {
			log.Println("close channel1")
		}()
		defer close(myChan)
		defer close(myChan2)
		defer close(doneChane)
		defer func() {
			log.Println("close channel2")
		}()

		wait.Wait()
	}()

	// 这两处代码等效，都可以实现从channel中读取数据，直到channel被关闭
	// for {
	// 	num, ok := <-myChan
	// 	if !ok {
	// 		log.Println("channel closed")
	// 		break
	// 	}
	// 	log.Println(num)
	// }

	// for num := range myChan {
	// 	log.Println(num)
	// }

	numList := []int{}
	nameList := []string{}

	var event = func() {
		for {
			select {
			case num, ok := <-myChan:
				if !ok {
					return
				}
				numList = append(numList, num)
			case name, ok := <-myChan2:
				if !ok {
					return
				}
				nameList = append(nameList, name)
			case <-doneChane:
				return

				// 协程超时处理
				// case <-time.After(1 * time.Second):
				// 	log.Println("timeout")
				// 	return
			}
		}
	}
	event()

	// 主线程结束 所有协程函数跟着结束
	log.Println("main function called", time.Since(startTime))
	log.Println("numList:", numList)
	log.Println("nameList:", nameList)
	// log.Printf("elapsed time: %s", time.Since(startTime))
}
