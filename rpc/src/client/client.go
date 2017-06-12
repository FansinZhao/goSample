package main

import (
	"log"
	"net/rpc"
	"server"
)

type Quo server.Quotient

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("连接异常", err)
	}
	args := &server.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("调用异常", err)
	}
	log.Println("输出结果:", reply)

	args1 := &server.Args{17, 8}
	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args1, &quotient, nil)
	replyCall := <-divCall.Done

	log.Println(replyCall.ServiceMethod, quotient)
}
