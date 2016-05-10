package server

import (
	"fmt"
	"net"
	"os"
)

//Server server instance
type Server struct {
	Host string
	Port string
	Type string
}

//New returns new server instance
func New(host string, port string, connType string) (server *Server) {
	return &Server{
		Host: host,
		Port: port,
		Type: connType,
	}
}

//Connect tests that we are inside
func (server *Server) Connect() (inc net.Listener) {
	fmt.Println("in test")
	inc, err := net.Listen(server.Type, server.Host+":"+server.Port)

	if nil != err {
		fmt.Println("Error while listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on " + server.Host + ":" + server.Port)

	return inc
}
