package main

import (
	"log"
	"sync"
	"time"
)

func setTimeOut(name string, wait *sync.WaitGroup) {
	log.Println("timeout-start,", name)
	time.Sleep(1 * time.Second)
	log.Println("timeout-end,", name)
	wait.Done()
}

func main() {
	var wait sync.WaitGroup
	log.Println("main function start")
	var startTime time.Time = time.Now()
	wait.Add(3)
	go setTimeOut("goroutine 1", &wait)
	go setTimeOut("goroutine 2", &wait)
	go setTimeOut("goroutine 3", &wait)

	wait.Wait()

	// 主线程结束 所有协程函数跟着结束
	log.Println("main function called", time.Since(startTime))
	// log.Printf("elapsed time: %s", time.Since(startTime))
}
