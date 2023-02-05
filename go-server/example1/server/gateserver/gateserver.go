package gateserver

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

func handleRequest(conn net.Conn) {
	// incoming request
	n := 0
	for {
		// Allocate header
		header := make([]byte, 2)

		// First read the two byte header
		_, err := io.ReadAtLeast(conn, header, 2)

		if err != nil {
			log.Fatal(err)
		}

		message_size := header[0]
		message_type := header[1]

		// Allocate for packet
		message_data := make([]byte, message_size)

		// And read the packet
		_, err = io.ReadFull(conn, message_data)

		fmt.Printf("%d: ", n)
		n++
		printReceivedBuffer(message_data, message_type)

		// Unmarshal binary data into Protocol Buffer gamepacket

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

func toGameMessage(buffer []byte, message_type byte) {
}

func StartServer() {

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
