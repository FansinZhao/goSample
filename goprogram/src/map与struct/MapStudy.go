// MapStudy
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Hello World!")
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2, "three": 3}
	mapCreate := make(map[string]float32)
	mapList = mapList
	mapList["one"] = 0
	mapCreate["key1"] = 2
	fmt.Println(mapList["one"])
	fmt.Println(mapList["two"])
	fmt.Println(mapList["four"])
	fmt.Println("key1", mapList["key1"])
	fmt.Println(mapCreate["key1"])
	fmt.Println(mapCreate["one"])
	//error make([2]int, 1, 3)
	fmt.Println(make([]int, 1, 3))
	//不能使用new,对于slice,map,channel使用make
	//	mm := new(map[string]int)
	//	mm["k"] = 1
	//	fmt.Println(mm["k"])
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 30 },
		//		1: func() int { return 10 },不允许重复key
	}
	fmt.Println(mf[2], mf[2]())
	ok := mf[1] //显示内存地址
	fmt.Println("ok", ok)
	_, ok1 := mf[11] //显示是否key存,为空是false
	fmt.Println("ok1", ok1)
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Found!")
	}
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "###.#")
	fmt.Println(str)
	fmt.Println(re.ReplaceAllStringFunc(searchIn, f))
}
