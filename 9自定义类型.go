package main

import (
	"log"
)

type Code int

const (
	Success Code = 1
	Failure Code = 0
)

func (c Code) getCodeMsg() string {
	switch c {
	case Success:
		return "success"
	case Failure:
		return "failure"
	default:
		return "unknown code"
	}
}

func server(name string) (code Code, msg string) {
	if name == "1" {
		return Success, Success.getCodeMsg()
	}
	return Failure, Failure.getCodeMsg()
}

func main() {
	code, msg := server("1")
	log.Println("code:", code, "msg:", msg)
}
