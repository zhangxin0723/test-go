package main

import (
	"io"
	"os"
)

func writeFile() {
	file, err := os.OpenFile("写入测试文件.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("文件打开失败")
	}
	defer file.Close()
	_, err = file.Write([]byte("hello world"))
	if err != nil {
		panic("文件写入失败")
	}
}

func copyFile() {
	srcFile, err := os.Open("写入测试文件.txt")
	if err != nil {
		panic("源文件打开失败")
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile("写入测试文件3.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic("目标文件打开失败")
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		panic("文件拷贝失败")
	}
}

func readDir() {
	files, err := os.ReadDir("./测试读取文件夹")
	if err != nil {
		panic("读取目录失败")
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			panic("获取文件信息失败")
		}
		println(file.Name(), info.Name())
	}
}

func main() {
	// err := os.WriteFile("写入测试文件1.txt", []byte("hello world1"), 0777)
	// if err != nil {
	// 	panic("文件写入失败")
	// }

	// writeFile()
	// copyFile()
	readDir()
}
