// string_package
package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "this is an example of a string "
	fmt.Printf("字符串:%s \n以this开头? %t 以G结尾 %t\n",
		str, strings.HasPrefix(str, "this"), strings.HasSuffix(str, "G"))
	fmt.Printf("包含an? %t\n", strings.Contains(str, "an"))
	fmt.Printf("包含an的位置? %d\n", strings.Index(str, "an"))
	fmt.Printf("包含an的最后位置? %d\n", strings.LastIndex(str, "an"))
	fmt.Printf("包含'中文'的最后位置? %d\n", strings.IndexRune("中文", '文'))    //注意使用'' 用来表示非ascii码
	fmt.Printf("包含'中文'的最后位置? %t\n", strings.ContainsRune("中文", '文')) //注意使用'' 用来表示非ascii码
	fmt.Println(strings.Replace("111222", "1", "a", 1))
	fmt.Println(strings.Replace("111222", "1", "a", 2))
	fmt.Println(strings.Count("111222", "1"))
	fmt.Println(strings.Repeat("hi go! ", 3))
	fmt.Println(strings.ToLower("Hi go! "))
	fmt.Println(strings.ToUpper("Hi go! "))
	fmt.Println(strings.TrimSpace("Hi go! "))
	fmt.Println(strings.Trim("Hi go! ", "Hi"))
	fmt.Println(strings.Fields("Hi go! "))
	fmt.Println(strings.Split("Hi go! ", " "))

	ss := "Go1|Go2|Go3"
	sl := strings.Split(ss, "|")
	for a, val := range sl {
		fmt.Printf("%d,%s-\n", a, val)
	}
	//	fmt.Println()
	s1 := strings.Join(sl, ";")
	fmt.Println(s1)

	fmt.Println(strconv.Atoi("66s66"))
	fmt.Println(strconv.Itoa(6666))
	fmt.Println(&ss)
	var i1 = 5
	var intP *int = &i1
	fmt.Println(i1, intP)
	var i2 = 5
	fmt.Println(&i2, *(&i2))
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6", k)
		fallthrough //不管k值如何变化,都会执行下一个分支,
	case 7:
		fmt.Println("was <= 7", k)
		fallthrough
	case 8:
		fmt.Println("was <= 8", k)
		fallthrough
	default:
		fmt.Println("default case", k)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	b()
	log.Println("loglog")
	log.Println("长度", len(str))
	aa := addFunc(1, 2, add)
	fmt.Println("fmt---", aa)

	s2 := "a中b文c在d这e里f"
	fmt.Println("原字符串", s2)

	for {
		index := strings.IndexFunc(s2, func(c rune) bool {
			if c > utf8.RuneSelf {
				s2 = strings.Replace(s2, string(c), " ", len(string(c)))
				return true
			}
			return false
		})
		if index == -1 {
			break
		}
	}
	log.SetFlags(log.Llongfile)
	log.Println("替换中文为空格后:", s2)
	log.Println(runtime.Caller(0))

}

func addFunc(a int, b int, f func(int, int) int) (c int) {
	c = f(a, b)
	return
}

func add(a int, b int) (c int) {
	c = a + b
	return
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) //先执行trace("b")得到值"b"
	fmt.Println("in b")
	a()
}
