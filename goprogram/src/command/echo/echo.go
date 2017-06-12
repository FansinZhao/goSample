package main

import (
	"flag"
	"os"
)

var NewLine = flag.Bool("n", false, "使用参数-实现换行打印!") //false 默认不开启

const (
	newLine = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	var str = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			str += " "
			if *NewLine {
				str += newLine
			}
		}
		str += flag.Arg(i)
	}
	os.Stdout.WriteString(str + "\n")
}
