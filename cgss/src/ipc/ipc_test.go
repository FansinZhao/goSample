package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{method, params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}
func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("echo", "From client1")
	resp2, _ := client1.Call("echo", "From client1")

	if resp1.Body != "From client1" || resp2.Body != "From client1" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
