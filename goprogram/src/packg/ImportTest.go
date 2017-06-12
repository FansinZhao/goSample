package packg

import (
	"fmt"
)

var ImportVar = "这是外部的变量"

//init 比需要引入的文件更早的初始化
func init() {
	fmt.Println("初始化外部文件!")
}
