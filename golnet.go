package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	// Connecting to server
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Reading initial server response
	reader := bufio.NewReader(conn)

	// Sending password
	fmt.Fprintf(conn, password+"\n")
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error during login:", err)
		os.Exit(1)
	}

	// Sending command
	fmt.Fprintf(conn, command+"\n")

	// Timeout waiting for a response
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// Waiting for command string
	expected := fmt.Sprintf("INF Executing command '%s' by Telnet", command)

	// Flag to check if the command string has been received
	responseReceived := false

	// Reading data from connection
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// If the socket is closed or there is an EOF error, we exit
			if strings.Contains(err.Error(), "closed") || strings.Contains(err.Error(), "EOF") {
				break
			}
			// If there is another error, we display it
			fmt.Println("Error reading from server:", err)
			break
		}

		// Printing received strings
		//fmt.Printf("Received line: %s", line)

		// If the string contains a command, mark the command as executed
		if strings.Contains(line, expected) {
			responseReceived = true
			break
		}
	}

	// If the command was executed
	if responseReceived {
		fmt.Println("Command was executed successfully")
		os.Exit(0)
	}

	// If time has passed and no response was received
	fmt.Println("Server is unresponsive. It might be frozen.")
	os.Exit(4)
}
