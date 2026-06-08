package main

import "log"

func main() {
	log.Println("123")
	arr := [3]int{1, 2, 3}
	log.Println(arr)

	arr1 := make([]int, 3)
	log.Println(arr1, len(arr1), cap(arr1))

	arr2 := []int{1, 2, 3}
	log.Println(arr2, len(arr2), cap(arr2))

	arr3 := []int{1, 2, 3}
	arr4 := arr3[0:2]
	log.Println(arr4, len(arr4), cap(arr4))
}
