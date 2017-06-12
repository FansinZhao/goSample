// FuncStudy
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var sum int
	func() {
		sum = 0.0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()
	fmt.Println(sum)

	hello := func() {
		fmt.Println("Hello World!")
	}
	hello()
	fmt.Printf("type %T,run %s\n", hello, hello)
	fmt.Println(f())
	end := time.Now()
	fmt.Println(end.Sub(start))

	a := [...]string{"a", "b", "c", "d"}
	for i, _ := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //从index=2开始,到index=5,共4个

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	slice2 := slice1[3:4]
	fmt.Println("slice2", len(slice2))
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
