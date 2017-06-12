package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	service := ":1200"
	tcp, err := net.ResolveTCPAddr("tcp4", service)
	checkErr(err)
	listener, err := net.ListenTCP("tcp4", tcp)
	checkErr(err)
	fmt.Println("建立服务器监听...")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	//超时 2min
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		n, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("读取到客户端信息....", string(request))
		if n == 0 {
			break
		} else if strings.TrimSpace(string(request[:n])) == "timestamp" {
			fmt.Println("timestamp")
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			fmt.Println("default")
			daytime := time.Now().String()
			fmt.Println(daytime)
			conn.Write([]byte(daytime))
		}
		conn.Close()
		request = make([]byte, 128)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal err %s", err.Error())
		os.Exit(-1)
	}
}
