package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	//	"golang.org/x/net/websocket"
)

func main() {
	service := ":1200"
	tcp, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	conn, err := net.DialTCP("tcp4", nil, tcp)
	checkErr(err)
	fmt.Println("连接成功!")
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	fmt.Println("发送成功!!", conn)
	//	b := make([]byte, 0, 1024)
	//	_, err = conn.Read(b)
	//	fmt.Println("返回:", string(b))
	b, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println("返回结果:", string(b))
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErr(err)
	fmt.Println("发送成功!!", conn)
	b, err = ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Println("返回结果:", string(b))
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stdin, "Fata error!%s", err.Error())
		os.Exit(-1)
	}
}
