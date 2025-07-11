package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        message, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Client disconnected:", conn.RemoteAddr())
            return
        }
        fmt.Printf("Received from %s: %s", conn.RemoteAddr(), message)
        conn.Write([]byte("Echo: " + message))
    }
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Server listening on :8080")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting:", err)
            continue
        }

        go handleConnection(conn) // 👉 spawn a new goroutine per client
    }
}