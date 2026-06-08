package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	// 一次性读取
	// // byteData, err := os.ReadFile("D:\\MY\\goLang\\test-go\\1.js")
	// byteData, err := os.ReadFile("1.js")
	// if err != nil {
	// 	log.Println("读取文件失败")
	// 	return
	// }
	// log.Println(string(byteData))

	// 分片读取
	// file, err := os.Open("1.js")
	// if err != nil {
	// 	log.Println("读取文件失败")
	// 	return
	// }
	// defer file.Close()
	// buf := make([]byte, 1024)
	// for {
	// 	n, err := file.Read(buf)
	// 	log.Println("n:", n)
	// 	if n == 0 || err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Println("读取文件失败")
	// 		return
	// 	}
	// 	log.Println(string(buf[:n]))
	// }

	// 缓冲读取
	file, err := os.Open("1.js")
	if err != nil {
		log.Println("读取文件失败")
		return
	}
	defer file.Close()

	// 按照行读取
	// reader := bufio.NewReader(file)
	// for {
	// 	line, _, err := reader.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	log.Println(string(line))
	// }

	reader := bufio.NewScanner(file)
	// 设置分隔符
	reader.Split(bufio.ScanWords)
	for reader.Scan() {
		log.Println(reader.Text())
	}
	if err := reader.Err(); err != nil {
		log.Println("读取文件失败")
	}
}
