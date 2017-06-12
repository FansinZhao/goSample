package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName  xml.Name `xml:"person"`
	Name     string   `xml:"name"`
	Sex      string   `xml:"sex,attr"`
	Age      int      `xml:"age"`
	WorkAddr WorkAddr `xml:"work"`
}

type WorkAddr struct {
	XMLName xml.Name `xml:"work"`
	No      string   `xml:"no,attr"`
	Addr    string   `xml:"addr"`
}

func main() {
	work := WorkAddr{No: "123"}
	p := &Person{Name: "赵锋", Age: 26, WorkAddr: work, Sex: "男"}

	b, _ := xml.MarshalIndent(p, "", "	")
	xmlStr := xml.Header + string(b)
	fmt.Println(xmlStr)
	var person Person
	err := xml.Unmarshal([]byte(xmlStr), &person)
	if err != nil {
		fmt.Println("error!")
	}
	fmt.Println(person.XMLName.Local, person.Sex, person.Name, person.Age, person.WorkAddr.No, person.WorkAddr.Addr)
}
