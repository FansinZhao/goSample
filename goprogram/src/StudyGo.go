//StudyGo
package main

import (
	"fmt"
	"net"
	"os"

	pk "./packg" //引入外部变量
)

func main() {
	fmt.Println("Hello Go World!")
	constStudy()
	varStudy()
	testFeilds()
	ip := net.ParseIP("172.20.0.193")
	fmt.Println(ip)
}

// init 这是一个保留函数名  默认会先执行这个init,再执行main
func init() {
	fmt.Println("init...最先执行")
	fmt.Println("打印引入包变量----" + pk.ImportVar)
}

const Pi float32 = 3.14159
const PI = 3.1415926
const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009

//constStudy 常量定义
func constStudy() {
	fmt.Println("常量定义-------------------")
	fmt.Println("隐式定义", Pi)
	fmt.Println("显式定义", PI)
	const hello = "Hello Go World!"
	fmt.Println(hello)
	//不会溢出,但是显示不全了?
	const ln2 = 0.123456789012345678979845564564876465456456476545645646545645654
	const log2e = 1 / ln2
	const billion = 1e9
	const hardEight = (1 << 100) >> 97
	fmt.Println("Ln2=", Ln2)
	fmt.Println("ln2=", ln2, " log2e=", log2e, " billion=", billion, " hardEight=", hardEight)
	//并发赋值
	const beef, two, c = "eat", 2, "veg"
	//	const Monday, Tuesday, WednesDay, Thursday, Friday, Saturday, Sunday = 1, 2, 3, 4, 5, 6, 7
	const (
		Monday, Tuesday, WednesDay         = 1, 2, 3
		Thursday, Friday, Saturday, Sunday = 4, 5, 6, 7
	)
	//iota 默认初始值为0,然后依次+1
	const (
		red = iota
		black
		white
	)
	fmt.Println(red, black, white)
	fmt.Println("常量定义-------------------")
}

var (
	a1   int
	b1   bool
	str1 string
)

//varStudy 变量学习
func varStudy() {
	fmt.Println("变量定义-------------------")
	var a int      //默认为0
	var b bool     //默认为false
	var str string //默认为空
	var p *int     //默认<nil>
	fmt.Println(a, b, str, p)
	var goos string = os.Getenv("GOOS")
	fmt.Println("你的操作系统是:" + goos)
	//简写
	path := os.Getenv("PATH")
	fmt.Println("PATH 是", path)
	fmt.Println("变量定义-------------------")
}

var a = "G"

//test 输出 G 0 G
func testFeilds() {
	fmt.Println("作用域范围.....")
	n()
	m()
	n()
}

func n() {
	fmt.Println(a)
}

func m() {
	a := 0
	fmt.Println(a)

}
