package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"server"
)

func main() {
	arti := new(server.Arith)
	rpc.Register(arti)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("监听器异常", e)
	}
	log.Println("监听服务启动....", 1234)
	http.Serve(l, nil)

}
