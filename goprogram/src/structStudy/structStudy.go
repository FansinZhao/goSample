// structStudy
package main

import (
	"fmt"
	"reflect"
)

type Foo map[string]string

type s struct {
	i1  int     "tag i1"
	f1  float32 "tag f1"
	str string  "tag str"
}

type inS struct {
	a int
	b int
}

type outS struct {
	x int
	y int
	int
	//	int一个类型只能有一个匿名内部类
	inS
}

func main() {
	s0 := new(s)
	s0.i1 = 1
	s0.f1 = 2.1
	(*s0).str = "3" //两种方式
	fmt.Println(s0)
	s1 := &s{11, 22.2, "333"}
	fmt.Println(s1)
	s3 := s{11, 22.2, "333"}
	fmt.Println(s3)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	b := make([]int, 5, 10) // len(b)=5, cap(b)=10
	fmt.Println(len(b), cap(b))
	//继续切片，注意len和cap的变化
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Println(len(b), cap(b))
	b = b[1:] // len(b)=4, cap(b)=4
	fmt.Println(len(b), cap(b))

	//	m := new(Foo)
	//	(*m)["a"] = 1
	//	fmt.Println(m)
	mm := make(Foo)
	mm["123"] = "abc"
	fmt.Println(mm)
	s2 := s{1, 2.2, "333"}
	ttType := reflect.TypeOf(s2)
	fmt.Println(ttType)
	ttField := ttType.Field(0)
	fmt.Println(ttField.Tag)
	outS := outS{11, 22, 33, inS{44, 55}}
	fmt.Println(outS)
	fmt.Println(outS.inS.AddThem())

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}

func (ins *inS) AddThem() int {
	return ins.a + ins.b
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
