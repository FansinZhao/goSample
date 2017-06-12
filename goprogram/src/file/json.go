package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type, City, Street string
}

type VCard struct {
	FirstName, LastName, Remark string
	Addresses                   []*Address
}

func main() {
	a1 := &Address{"private", "深圳", "草围"}
	a2 := new(Address)
	a2.Type = "work"
	a2.City = "sz"
	a2.Street = "科技园"
	card := VCard{"feng", "zhao", "it", []*Address{a1, a2}}
	fmt.Println(card)
	js, _ := json.Marshal(card)
	fmt.Printf("%s\n", js)
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(card)
	if err != nil {
		log.Println("json写入文件失败!")
	}
	var c VCard
	er := json.Unmarshal(js, &c)
	if er != nil {
		fmt.Println(c)
	}

	defer func() {
		log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
		log.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start")
	panic("崩溃了")

	var notErr = errors.New("没有发现!")
	fmt.Println(notErr)

}
