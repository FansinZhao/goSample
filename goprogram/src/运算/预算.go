package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//main 居然可以使用中文文件夹和文件名 不建议使用,仅学习使用
func main() {
	fmt.Println("中文可以么?")
	var b = 10
	fmt.Println((b == 5))              //false
	fmt.Println((b == 10))             //true
	fmt.Println((b == 10) || (b == 5)) //true
	fmt.Println((b == 10) && (b == 5)) //false
	fmt.Println(!(b == 10))            //false
	var a int
	fmt.Println(reflect.TypeOf(a), a)
	a = a + a //可以执行
	fmt.Println(reflect.TypeOf(a), a)
	var c int32
	//	c = a 不能执行
	c = int32(a)
	var d int64
	//	d = a 不能执行
	d = int64(a)
	fmt.Println(reflect.TypeOf(d), d)
	//不能正常执行
	//	c = a + a 异常
	//c = a + 15 异常
	c = c + 15
	fmt.Println(reflect.TypeOf(c), c)
	//	c / 0
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /", a)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		a := rand.Intn(8)
		fmt.Printf("%d /", a)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}

}
