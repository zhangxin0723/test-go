package main

import (
	"log"
)

func add(a, b int) int {
	return a + b
}

func add2(numList ...int) int {
	sum := 0
	for _, num := range numList {
		sum += num
	}
	return sum
}

func r1(a, b int) (val int, ok bool) {
	if a > b {
		val = a
		ok = true
		return
	} else {
		return
	}
}

func setName(name *string) {
	*name = "李四"
}

func init() {
	log.Println("init function called")
}

func init() {
	log.Println("init function called again")
}

func main() {
	// log.Println("111")
	// log.Println(add(1, 2))
	// log.Println(add2(1, 2, 3, 4))
	// result, ok := r1(3, 2)
	// if ok {
	// 	log.Printf("result: %d", result)
	// } else {
	// 	log.Println("a is not greater than b")
	// }
	defer func() {
		log.Println("defer function called")
	}()

	var name string = "张三"
	log.Printf("name: %s", name)
	setName(&name)
	log.Printf("name: %s", name)
}
