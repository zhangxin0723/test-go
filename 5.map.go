package main

import "log"

func main() {
	userMap := map[int]string{
		1: "张三",
		2: "李四",
		3: "王五",
	}
	value, ok := userMap[4]
	log.Println(value, ok)
	log.Printf("%#v", userMap[4])
	log.Println(userMap)

	userMap[4] = "赵六"
	log.Println(userMap)

	userMap[1] = "张三丰"
	log.Println(userMap)

	delete(userMap, 4)
	log.Println(userMap)
}
