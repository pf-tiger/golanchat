// this file is a server application:
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
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
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("### Closed connection from", username)
			break
		}
		message = strings.TrimSpace(message)
		fmt.Println(username, "at", time.Now().Format("15:04:05"), ">", message)
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
