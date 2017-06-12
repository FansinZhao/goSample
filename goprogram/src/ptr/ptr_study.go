// ptr_study
package main

import (
	"fmt"
)

func main() {
	var str = "aaaa"
	fmt.Println(&str)
	var i1 = 5
	var intP *int = &i1
	fmt.Println(i1, intP)
	var i2 = 5
	fmt.Println(&i2, *(&i2))
	var strP *string = &str
	*strP = "bbbb" //指针可以修改原值
	fmt.Println(str)
	str = "ccc"
	fmt.Println(str)
	var p *int = nil
	*p = 0

}
