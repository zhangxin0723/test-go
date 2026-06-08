package main

import (
	"encoding/json"
	"log"
)

type User struct {
	Name     string `json:"name" db:"name"`
	Age      int    `json:"age" db:"age"`
	Password string `json:"-" db:"password"`
}

func main() {
	user := User{Name: "张三", Age: 30, Password: "secret"}
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}
	log.Println(string(data))
}
