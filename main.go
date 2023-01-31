package main

// takes an argument for the target server.
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run client.go [server IP address]")
		os.Exit(1)
	}

	serverAddr := os.Args[1] + ":8080"
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type 'exit' to quit.")
	for {
		reader := bufio.NewReader(os.Stdin)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fmt.Print("> ", timestamp, ":")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		timestamp = time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintf(conn, "%s: %s\n", timestamp, text)
	}
	fmt.Println("Disconnected from server.")
}
