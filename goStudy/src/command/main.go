// command project main.go
package main

import (
	"fmt"

	"command/printmsg" //从src下开始找
)

func main() {
	fmt.Println("Hello World!")
	printmsg.PrintMsg()
}
