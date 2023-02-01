// package main

// // takes an argument for the address/hostname of the messge server

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

const displayLimit = 30
const TitleBanner = "golanchat by pf_tiger"

func main() {
	arguments := os.Args
	// when no address/hostname is specified
	if len(arguments) == 1 {
		fmt.Println("Please provide a server address/hostname")
		return
	}

	conn, err := net.Dial("tcp4", arguments[1]+":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to the chat server:", arguments[1], "Type messages and press Enter to send.")

	numlines := 0
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("You > ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		numlines++

		currentTime := time.Now().Format("15:04:05")

		message := currentTime + ":" + text

		// command "exit!" to close connection from the server
		if text == "exit!\n" {
			fmt.Println("Closing connection to server")
			conn.Close()
			break
		}

		fmt.Print("\033[A")  // move cursor up
		fmt.Print("\033[2K") // clear current line
		// fmt.Print("\033[A")  // move cursor up
		// fmt.Print("\033[2K") // clear current line

		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Connection closed")
			break
		}
		fmt.Println(currentTime, " You > ", text)

		// cleaning the UI
		if numlines >= displayLimit {
			// write codes to clean the UI
			for i := 0; i < numlines; i++ {
				fmt.Print("\033[A") // move cursor up
			}
			fmt.Print("\033[2K") // clear current line
			for i := 0; i < numlines; i++ {
				fmt.Print("\033[B") // move cursor up
			}
		}
	}
}

// // function to receive and display messages from the server/other clients

// // further functions for switching debugging modes is wanted
