package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i float64 = 5
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(&i) //将指针传递过去,修改后,原值就可以改变了
	fmt.Println(v.CanSet(), i)
	fmt.Println(v.Elem().CanSet())
	fmt.Println(t.Kind())
}
