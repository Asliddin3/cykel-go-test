package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "143.42.61.34"
	CONN_PORT = "9679"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	client, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	conn, err := l.Accept()
	fmt.Println(err)
	res := addByte([]byte("*CMDS,OM,860537062636022,200318123020,L0,0,1234,1497689816#\n"))
	_, err = conn.Write([]byte(res))
	var resultTemp []byte
	fmt.Println("read err", err)
	for len(resultTemp) == 0 {
		_, err = client.Read(resultTemp)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("read error", err, string(resultTemp))
	conn.Close()
	conn1, err := l.Accept()
	fmt.Println("answer error", err)
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn1.Read(buf)
	conn1.Close()
	fmt.Println(reqLen)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(string(buf))
	conn2, err := l.Accept()

	if err != nil {
		fmt.Println("some error accepting from lock", err)
	}
	res1 := addByte([]byte("*CMDS,OM,860537062636022,200318123020,Re,L0#\n"))
	fmt.Println(res1)
	conn2.Write([]byte(res1))
	conn2.Close()
	resultTemp = []byte{}
	client.Read(resultTemp)
	for len(resultTemp) == 0 {
		_, err = client.Read(resultTemp)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("client read result", resultTemp)
	// 	for {
	// 		// Listen for an incoming connection.
	// 		conn, err := l.Accept()
	// 		if err != nil {
	// 			fmt.Println("Error accepting: ", err.Error())
	// 			os.Exit(1)
	// 		}
	// 		// Handle connections in a new goroutine.
	// 		// go handleRequest(conn)
	// 	}
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
	reqLen, err := conn.Read(buf)
	fmt.Println(reqLen)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println(string(buf))
	// Send a response back to person contacting us.
	conn.Write([]byte("*CMDS ,OM,860537062636022,000000000000,L0,0,1234,1497689816#\n"))
	// Close the connection when you're done with it.
	conn.Close()
}
