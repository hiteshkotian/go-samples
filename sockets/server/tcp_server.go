package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

const defaultPort = "7777"
const defaultTCPNetworkType = "tcp"

func main() {

	port := flag.String("port", defaultPort, "port server will listen on")
	tcp := flag.String("networkType", defaultTCPNetworkType, "TCP Network type. Supported values are\n\ttcp\n\ttcp4\n\ttcp6\r\n")
	flag.Parse()

	service := fmt.Sprintf("localhost:%s", *port)
	fmt.Println("Server listening on: ", service)
	tcpAddr, err := net.ResolveTCPAddr(*tcp, service)
	checkError(err)

	listener, err := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "TCP Server error : %s", err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
