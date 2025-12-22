package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/gorilla/websocket"
)

const (
    Char_reset = "\033[0m"
    Color_cyan = "\033[36m"
    Color_green = "\033[32m"
)

func main() {
	var serverUrl string
	fmt.Print("Server url : ")
	fmt.Scanln(&serverUrl)

	StartClient(serverUrl)
}

func StartClient(serverUrl string) {
	dialer := websocket.DefaultDialer
	dialer.NetDialContext = (&net.Dialer{
		Resolver: &net.Resolver{
			PreferGo:     false,
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

	fmt.Println("Connection established!")

	rl, err := readline.New(Color_green + "> " + Char_reset)
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	log.SetOutput(rl.Stderr())

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				os.Exit(0)
			}
			fmt.Fprintf(rl.Stdout(), Color_cyan + "%s\n" + Char_reset, string(message))
		}
	}()

	for {
		text, err := rl.Readline()
		if err != nil { 
			break
		}

		if strings.TrimSpace(text) == "" {
			continue
		}

		if text == "exit" {
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			break
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
		
	}
}
