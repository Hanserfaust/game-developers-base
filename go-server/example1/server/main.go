package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "7777"
	TYPE = "tcp"
)

func main() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Server up on", HOST, ":", PORT)

	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

type GameMessage struct {
	name      string
	client_id string
	data      []int32
}

func handleRequest(conn net.Conn) {
	// incoming request
	n := 0
	for {
		header := make([]byte, 2)

		// First read the two byte header
		_, err := io.ReadAtLeast(conn, header, 2)

		if err != nil {
			log.Fatal(err)
		}

		message_size := header[0]
		message_type := header[1]

		message_data := make([]byte, message_size)

		// And read the packet
		_, err = io.ReadFull(conn, message_data)

		fmt.Printf("%d\n", n)
		n++
		printReceivedBuffer(message_data, message_type)
		// fmt.Println(game_message.name)
	}

	// write data to response
	// time := time.Now().Format(time.ANSIC)
	// responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	// conn.Write([]byte(responseStr))

	// close conn
	// conn.Close()
}

func printReceivedBuffer(buffer []byte, message_type byte) {
	fmt.Printf("Got message of type 0x%x:", message_type)

	encodedStr := hex.EncodeToString(buffer)
	fmt.Printf("%s\n", encodedStr)
}
