package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "time"
    "net"
	"github.com/gorilla/websocket"
)

func StartClient(serverUrl string) {
    dialer := websocket.DefaultDialer
    dialer.NetDialContext = (&net.Dialer{
        Resolver: &net.Resolver{
            PreferGo: false, 
            StrictErrors: false,
        },
        Timeout:   10 * time.Second,
        KeepAlive: 30 * time.Second,
    }).DialContext

	serverURL := "wss://" + serverUrl
	fmt.Printf("Connecting to server... \n")

	conn, _, err := dialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected established!")

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("\nDisconnected from server.")
				os.Exit(0)
			}
			fmt.Printf("\r>%s\n> ", string(message))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ") 
	
	for scanner.Scan() {
		text := scanner.Text()

		if text == "exit" {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			break
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}

		fmt.Print("> ")
	}
}

func main() {
    var serverUrl string
    fmt.Print("Server url : ")
    fmt.Scanln(&serverUrl)

	StartClient(serverUrl)
}
