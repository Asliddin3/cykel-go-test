package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "9679"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	fmt.Println(err)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("some error", err)
			break
		}
		go handleRequest(conn)
		conn.Close()
	}
}

func addByte(b2 []byte) []byte {
	arrByte := make([]byte, 2)
	arrByte[0] = 0xFF
	arrByte[1] = 0xFF
	arrByte = append(arrByte, b2...)
	return arrByte
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("read result", string(buf))
	// Send a response back to person contacting us.
	res := addByte([]byte("*CMDS,OM,860537062636022,200318123020,L0,0,1234,1497689816#\n"))
	fmt.Println("send message",string(res))
	conn.Write([]byte(res))
	// Close the connection when you're done with it.
	conn.Close()
}
