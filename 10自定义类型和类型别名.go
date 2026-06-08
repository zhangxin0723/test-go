package main

import "log"

// 自定义类型
type MyType int

// 类型别名
type MyType2 = int

// 区别
// 1.类型别名不会创建新的类型，它只是为现有类型创建了一个别名，而自定义类型会创建一个新的类型。
// 2.自定义类型可以有自己的方法，而类型别名不能有方法，因为它只是现有类型的别名。
// 3.自定义类型可以用于接口实现，而类型别名不能用于接口实现，因为它只是现有类型的别名。
// 4.自定义类型可以用于类型断言，而类型别名不能用于类型断言，因为它只是现有类型的别名。
// 5.自定义类型可以用于类型转换，而类型别名不能用于类型转换，因为它只是现有类型的别名。
// 6.自定义类型可以用于类型检查，而类型别名不能用于类型检查，因为它只是现有类型的别名。
// 7.自定义类型可以用于类型推断，而类型别名不能用于类型推断，因为它只是现有类型的别名。
// 8.自定义类型可以用于类型推导，而类型别名不能用于类型推导，因为它只是现有类型的别名。
// 9.自定义类型可以用于类型约束，而类型别名不能用于类型约束，因为它只是现有类型的别名。
// 10.自定义类型可以用于类型别名，而类型别名不能用于自定义类型，因为它只是现有类型的别名。

// 类型别名和自定义类型的区别总结
// 1、不能绑定方法
// 2、打印类型还是原始类型
// 3、和原始类型比较，类型别名不用转换

const code1 MyType = 1

func (m MyType) getCodeMsg() string {
	return "code1"
}

const code2 MyType2 = 1

// func (m MyType2) getCodeMsg() string {

// }

func main() {
	log.Printf("%T", code1) // main.MyType
	log.Printf("%T", code2) // int

	age := 1
	// if code1 == age { // 不能直接比较，因为code1是MyType类型，而age是int类型
	// if code1 == MyType(age) {
	if int(code1) == age {
		log.Println("code1 == 1")
	}

	if code2 == age {
		log.Println("code2 == 1")
	}
}
