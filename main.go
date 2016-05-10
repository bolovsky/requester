package main

import (
	"bytes"
	"fmt"
	"net"
	"os"

	"github.com/bolovsky/requester/Server"
	"github.com/bolovsky/requester/WebWorker"
)

/**
  connHost server host
  connPort server port
  connType server type (eg: tcp)
*/
const (
	connHost = "localhost"
	connPort = "2345"
	connType = "tcp"
)

var pool *WebWorker.Pool

func main() {
	fmt.Printf("Starting Server...\n")

	pool = WebWorker.NewPool(5)

	// do requests ad eternum or until interrupt
	openSocketServer()

	pool.ShutDown()

	fmt.Printf("Stopping Server...\n")
}

func openSocketServer() {
	server := server.New(connHost, connPort, connType)
	inc := server.Connect()

	//make sure the listener closes when the application is closing
	defer inc.Close()

	for {
		conn, err := inc.Accept()
		if nil != err {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleMessageReceived(conn)
	}
}

func handleMessageReceived(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	lenght, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	//show what we got
	if lenght > 0 {
		//convert buffer to string
		dataStr := string(buf[:bytes.IndexByte(buf, 0)])
		pool.QueueJob(WebWorker.Job{
			Payload: dataStr,
			URL:     "http://localhost:8080",
		})
	}

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received\n"))
	// Close the connection when you're done with it.
	conn.Close()
}
