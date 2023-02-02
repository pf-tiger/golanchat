// this file is a server application:
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp4", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening for connections...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

var connections []net.Conn

// receive and display messages
func handleConnection(conn net.Conn) {
	username := conn.RemoteAddr().String()
	connections = append(connections, conn)
	fmt.Println("### Accepted connection from", username)
	for {
		full_message, err := bufio.NewReader(conn).ReadString('\n')

		// split and assign fisrt 8 letters for timestamp, and the message
		// currently used format for time: HH:MM:SS

		// add if statements to handle length errors
		// current error for full_message not having sufficient length panic: runtime error: index out of range [1] with length 1
		splitMessage := strings.SplitAfterN(full_message, ":", 4)
		timeinfo := splitMessage[0] + ":" + splitMessage[1] + ":" + splitMessage[2]
		message := splitMessage[3]

		if err != nil {
			fmt.Println("### Closed connection from", username)
			break
		}
		message = strings.TrimSpace(message)
		fmt.Println(username, "at", timeinfo, ">", message)
		for _, connection := range connections {
			if connection != conn {
				connection.Write([]byte(username + " > " + message + "\n"))
			}
		}
	}
	conn.Close()
	for i, connection := range connections {
		if connection == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

// Send a message to all clients
// func sendMessageToAllClients(message string) {
// 	for _, client := range connections {
// 		_, err := client.Write([]byte(message))
// 		if err != nil {
// 			log.Print("Error sending message:", err)
// 		}
// 	}
// }
