![License](https://img.shields.io/github/license/cocosilex/ws-chat)
![Go](https://img.shields.io/badge/Go-Client-blue)
![Java](https://img.shields.io/badge/Java-Server-orange)

# Websocket Chat
This project include a java websocket server and a go client designed to communicate with the server.

# How to build/use
## Client :
You will need the [go](https://go.dev/) CLI, you can either compile and run it after or directly run by doing `go run main.go`. After you'll need to enter the url of the server you want to connect.

## Server :
Option 1 : use [docker](https://www.docker.com/), simply run `docker compose up`, it will build/start automatically a container
Option 2 : use [gradle](https://gradle.org/), you will also need [java 21](https://adoptium.net/temurin/releases?version=21&os=any&arch=any) installed on your machine. Run `gradle build run` to build and start the script.
