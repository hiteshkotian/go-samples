package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// Signal for the program to know that all the
// requests have been processed.
const closeSignal = "close"

const bufferSize = 10

// result structure representing request of a client call.
type result struct {
	requestID string
	result    string
}

// printResult method prints the result object properties
func (r *result) printResult() {
	fmt.Printf("[%s] : %s\n", r.requestID, r.result)
}

// makeClientCall makes a call to the tcp server.
func makeClientCall(service, name string, resultChannel chan result) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	serverResponse := readServerResponse(conn)
	resultChannel <- result{name, serverResponse}
}

// readServerResponse reads the server response from the
// TCPConn output stream and returns the string response
func readServerResponse(conn *net.TCPConn) string {
	reply := make([]byte, bufferSize)
	var bytesRead int
	var serverResponse string
	var err error
	for {
		bytesRead, err = conn.Read(reply)

		if bytesRead <= 0 {
			break
		}
		checkError(err)

		serverResponse += string(reply)
	}
	serverResponse += "\n"
	return serverResponse
}

// createClientCalls creates go routines to make server calls. The number of goroutines
// to be created are specified by the numOfThreads argument.
func createClientCalls(server string, numOfThreads int, resultChannel chan result) {
	go func() {
		for i := 0; i < numOfThreads; i++ {
			makeClientCall(server, fmt.Sprintf("%d", i), resultChannel)
		}
	}()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	server := os.Args[1]
	fmt.Println("Server address: ", server)

	resultChannel := make(chan result, 3)
	createClientCalls(server, 10, resultChannel)

	var resultObject result
	for {
		select {
		case resultObject = <-resultChannel:
			resultObject.printResult()
		// Set a timeout of 3 seconds. If you don't get any result from
		// the resultChannel gracefully close the program
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout occurred...ending client")
			os.Exit(0)
		}
	}
}

// checkError checks the error object and prints out the error message before exiting
// in case of any errors
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		panic(err)
	}
}
