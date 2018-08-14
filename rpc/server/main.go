package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloService struct {
	*rpc.Client
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func (p HelloService) Hello(request string, replay *string) error {
	*replay = "hello: " + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", HelloService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen tcp error: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error: ", err)
	}
	rpc.ServeConn(conn)
}
