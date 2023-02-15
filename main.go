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
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := l.Accept()
		fmt.Println("accepted error", err)
		if err != nil {
			fmt.Println("some error", err)
			break
		}
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	buf := make([]byte, 2048)

	lenBuf, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("read result", string(buf), "with len", lenBuf)

	res := addByte([]byte("*CMDS,OM,860537062636022,0000000000,L0,0,1234,1497689816#\r\n"))
	fmt.Println("send message", string(res))
	_, err = conn.Write([]byte(res))
	fmt.Println("write error", err)
	conn.Close()
}
func addByte(b2 []byte) []byte {
	arrByte := make([]byte, 2)
	arrByte[0] = 0xFF
	arrByte[1] = 0xFF
	arrByte = append(arrByte, b2...)
	return arrByte
}
