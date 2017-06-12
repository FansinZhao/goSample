package main

import (
	"fmt"
)

type MyList struct {
}

type Any interface{}

func main() {
	var lst List
	//	CountInto(lst, 1, 10)
	if LongEnough(lst) {
		fmt.Println("够长了....")
	}
	plist := new(List)
	CountInto(plist, 1, 10)
	if LongEnough(plist) {
		fmt.Println("不够了....")
	}
	//	ml := new(MyList)
	//	fmt.Println(ml.(Any))
	aa := [5]int{1, 2, 3, 4, 5}
	as := aa[:3]
	fmt.Println(as)
	// 先定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]
	fmt.Println(mySlice)

}

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func (l List) Len() int {
	return len(l)
}

func CountInto(a Appender, start, end int) {
	fmt.Println(start, end)
	for i := start; i <= end; i++ {
		a.Append(i)
	}
	fmt.Println(a)
}

func LongEnough(l Lener) bool {
	fmt.Println(l.Len())
	return l.Len()*10 > 42
}
