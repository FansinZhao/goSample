package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string([]rune(inputString)[2:5]) + "\r\n"
		fmt.Println(outputFile)
		_, err := outputWriter.WriteString(outputString)
		outputWriter.Flush() //不能遗漏
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	fmt.Println("Conversion done")
}
