package main

import (
	"errors"
	"log"
	"runtime/debug"
)

// func init() {
// 	// 阻止main函数的执行，直接抛出一个错误
// 	panic("init panic")
// }

func div(a, b int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
			log.Println("Stack trace:", string(debug.Stack()))
		}
	}()
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := div(10, 0)
	if err != nil {
		log.Println("Error:", err)
	} else {
		log.Println("Result:", result)
	}
}
