package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	bubble "sorter/algorithms/bubblesort"
	quick "sorter/algorithms/quicksort"
)

var infile *string = flag.String("i", "", "需要排序的文件")
var outfile *string = flag.String("o", "", "保存已排序的文件")
var algorithm *string = flag.String("a", "qsort", "排序算法")

func readValues(infile string) (values []int, err error) {
	file, e := os.Open(infile)
	if e != nil {
		fmt.Println("找不到排序文件", infile)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, e := reader.ReadLine() //查询出来不含'\n'或'\r\n'
		if e != nil {
			if e != io.EOF {
				err = e
				fmt.Println("读取文件错误!")
			}
			break
		}

		if isPrefix {
			fmt.Println("单行记录太长了")
			return
		}

		i, e1 := strconv.Atoi(string(line))
		if e1 != nil {
			fmt.Println("无法转换数值", line, e1)
		}
		values = append(values, i)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.OpenFile(outfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建文件失败!", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range values {
		vs := strconv.Itoa(v)
		writer.WriteString(vs + "\n")
	}
	writer.Flush()
	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println(values)
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			quick.QuickSort(values)
		case "bubblesort":
			bubble.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}

	//	bubble.BubbleSort(values)
	//	quick.QuickSort(values)
}
