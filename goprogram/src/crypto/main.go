// main.go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	Teststr := "hello"
	hash := md5.New()
	hash.Write([]byte(Teststr))
	b := hash.Sum([]byte(""))
	fmt.Printf("%x\n", b)
}
