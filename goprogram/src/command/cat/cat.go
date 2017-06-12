package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var LineNumber = flag.Bool("n", false, "使用-n显示文件行号")

func main() {
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		defer f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))

	}
}

func cat(r *bufio.Reader) {
	var n = 0
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *LineNumber {
			n++
			fmt.Fprintf(os.Stdout, "%d %s", n, buf)
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}
