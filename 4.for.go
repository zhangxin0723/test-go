package main

import (
	"log"
)

func main() {
	// var age int = int(0)
	// for i := 0; i < 10; i++ {
	// 	// println(i)
	// 	age += i
	// }
	// log.Println(age)

	list := []string{"a", "b", "c", "d", "e"}
	var list1 []string = []string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(list); i++ {
		log.Println(list[i])
	}
	log.Println(list1)
	for index, value := range list {
		log.Printf("index: %d, value: %s", index, value)
	}

	userMap := map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}

	for key, value := range userMap {
		log.Printf("key: %d, value: %s", key, value)
	}
}
