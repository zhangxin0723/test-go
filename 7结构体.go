package main

// 定义一个函数类型
type MyFunc func(int) int

type Parent struct {
	name string
}

// 定义一个结构体类型
type MyStruct struct {
	Parent
	name string
	age  int
}

// 定义一个接口类型
type MyInterface interface {
	SayHello() string
}

func (s *MyStruct) SayHello(name string) string {
	s.name = name
	return "Hello, my name is " + s.name + " and you are " + name + "." + " I am " + string(s.age) + " years old." + " My parent name is " + s.Parent.name + "."
}

//  结构体tag
type User struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

func main() {
	// // 使用结构体
	// parent := Parent{name: "父亲"}
	// person := MyStruct{Parent: parent, name: "张三", age: 30}
	// println(person.SayHello("张三2"))

	// // 使用函数类型
	// var f MyFunc = func(x int) int {
	// 	return x * x
	// }
	// println(f(5))

}
