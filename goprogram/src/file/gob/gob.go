package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3, 4, 5, "O(∩_∩)O哈哈~"})
	if err != nil {
		log.Fatal("编码错误", err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("解密错误", err)
	}
	fmt.Printf("%d %d %q\n", *q.X, *q.Y, q.Name) //q是个指针

	if f, err := 1 / 0; err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
