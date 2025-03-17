package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: ./golnet <ip> <port> <password> <command>")
		os.Exit(1)
	}

	ip := os.Args[1]
	port := os.Args[2]
	password := os.Args[3]
	command := os.Args[4]

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Reading initial server response
	reader := bufio.NewReader(conn)
	_, _ = reader.ReadString('\n')

	// Sending password
	fmt.Fprintf(conn, password+"\n")
	time.Sleep(500 * time.Millisecond)

	// Sending command
	fmt.Fprintf(conn, command+"\n")
	time.Sleep(500 * time.Millisecond)

	// Sending exit command
	fmt.Fprintf(conn, "exit\n")
	time.Sleep(500 * time.Millisecond)
}
