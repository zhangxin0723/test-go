package main

import (
	"log"
	"sync"
)

//sync.Map
//sync.Map 是 Go 语言标准库中的一个并发安全的 map 类型，它提供了一些方法来操作 map，包括 Store、Load、Delete、Range 等。
//sync.Map 的设计目标是提供一种简单、高效的方式来操作并发 map，而不需要使用互斥锁或其他同步机制来保护 map 的访问。

// // 线程不安全
// var sum int
// var wait sync.WaitGroup

// func increase() {
// 	for i := 0; i < 100; i++ {
// 		sum++
// 	}
// 	wait.Done()
// }

// func reduce() {
// 	for i := 0; i < 100; i++ {
// 		sum--
// 	}
// 	wait.Done()
// }

// func main() {
// 	wait.Add(2)
// 	go increase()
// 	go reduce()
// 	wait.Wait()
// 	log.Println(sum)
// }

// 线程安全  使用lock
// var sum int
// var wait sync.WaitGroup
// var lock sync.Mutex

// func increase() {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	for i := 0; i < 100; i++ {
// 		sum++
// 	}
// 	wait.Done()
// }

// func reduce() {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	for i := 0; i < 100; i++ {
// 		sum--
// 	}
// 	wait.Done()
// }

// func main() {
// 	wait.Add(2)
// 	go increase()
// 	go reduce()
// 	wait.Wait()
// 	log.Println(sum)
// }

// 线程不安全的map
// var myMap = make(map[string]int)

// func main() {

// 	go func() {
// 		for {
// 			myMap["key1"] = 1
// 		}
// 	}()

// 	go func() {
// 		for {
// 			log.Println(myMap["key1"])
// 		}
// 	}()

// 	select {}

// }

// 线程安全的map
func main() {
	var myMap = sync.Map{}
	myMap.Store("key1", 1)
	myMap.Store("key2", 2)

	val, ok := myMap.Load("key1")
	if ok {
		log.Println(val)
	}

}
