package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    reader := bufio.NewReader(os.Stdin)
    serverReader := bufio.NewReader(conn)

    for {
        fmt.Print("Enter message: ")
        text, _ := reader.ReadString('\n')
        conn.Write([]byte(text))

        reply, _ := serverReader.ReadString('\n')
        fmt.Println("Server replied:", reply)
    }
}