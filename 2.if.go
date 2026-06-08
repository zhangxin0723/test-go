package main

func main() {
	age := 0
	if age < 18 {
		println("未成年")
	}

	switch age {
	case 0:
		println("年龄不能为负数")
	case 1:
		println("未成年")
	default:
		println("成年人")
	}
}
