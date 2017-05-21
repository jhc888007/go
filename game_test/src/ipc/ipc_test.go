package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	return &Response{method, "Echo:" + params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, err1 := client1.Call("A", "From Client1")
	resp2, err2 := client2.Call("B", "From Client2")

	if resp1.Body != "Echo:From Client1" || resp2.Body != "Echo:From Client2" ||
		err1 != nil || err2 != nil {
		t.Error("IpcClient.Call failed. resp1:", resp1, ", resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
