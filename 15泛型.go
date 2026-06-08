package main

import (
	"encoding/json"
	"log"
)

func plus(a, b int) int {
	return a + b
}

func plus1[T int | float64](a, b T) T {
	return a + b
}

type User[T any] struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Info T      `json:"info,omitempty"`
}

type UserInfo struct {
	Address string `json:"address"`
}

func main() {
	println(plus(1, 2))
	println(plus1(1, 2.3)) // 编译错误
	log.Printf("%.2f", plus1(1, 2.3))

	user := User[UserInfo]{
		Name: "Alice",
		Age:  30,
		Info: UserInfo{Address: "Beijing"},
	}
	data, _ := json.Marshal(user)
	log.Println(string(data))

	var info User[UserInfo]
	json.Unmarshal([]byte(`{"name":"Alice","age":30,"info":{"address":"Beijing"}}`), &info)
	log.Println(info.Info.Address)
}
