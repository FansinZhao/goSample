package main

import (
	"fmt"
)

var a = "G"

func main() {
	n()
	m()
	n()
	str := ifStudy()
	fmt.Println(str)
}

func ifStudy() string {
	var i = 1
	if i == 1 {
		return "1"
	} else if i == 2 {
		return "2"
	} else if i == 3 {
		return "3"
	} else {

		return "0"
	}
	return "-1"
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
