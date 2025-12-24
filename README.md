![License](https://img.shields.io/github/license/cocosilex/ws-chat)
![Go](https://img.shields.io/badge/Go-Client-blue)
![Java](https://img.shields.io/badge/Java-Server-orange)

# Websocket Chat
This project includes a Java websocket server and a Go client designed to communicate with the server.

## Client :
- Requires the [Go](https://go.dev) CLI.
- Go to the client directory `cd client`.
- Start with `go run main.go`.
- Enter the URL of the server you want to connect to.

## Server :
- Requires [Docker](https://www.docker.com) **OR** [Gradle](https://gradle.org) + [Java 21](https://adoptium.net/temurin/releases?version=21&os=any&arch=any).

- Navigate to the server directory `cd server`.
- For Docker, make sure you have the docker compose CLI and simply run `docker compose up`.
- For Gradle/Java 21, run `gradle build run` and it will build/start the server.
