package main

// takes an argument for the address/hostname of the messge server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args

	// when no address/hostname is specified
	if len(arguments) == 1 {
		fmt.Println("Please provide a server address/hostname")
		return
	}

	conn, err := net.Dial("tcp4", arguments[1]+":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the chat server:", arguments[1], "Type messages and press Enter to send.")

	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				break
			}
			message = strings.TrimSpace(message)
			fmt.Println(message)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}
		message = message + "\n"
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		if message == "exit!\n" {
			fmt.Println("Closing connection to server")
			defer conn.Close()
			break
		}
		fmt.Print(time.Now().Format("15:04:05"), " You > ", message)
	}
}

// function to receive and display messages from the server/other clients

// further functions for switching debugging modes is wanted
