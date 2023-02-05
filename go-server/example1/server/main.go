package main

import (
	"fmt"
	"io"
	"log"
	"messages.pb"
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
	fmt.Printf("Server up on %v:%v", HOST, PORT)

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
	for {
		buf_len := make([]byte, 1)
		message_data := make([]byte, 256)

		_, err := io.ReadAtLeast(conn, message_data, int(buf_len[0]))

		if err != nil {
			log.Fatal(err)
		}
		var game_message = decodeBuffer(buffer)
		fmt.Println(game_message.name)
	}

	// write data to response
	// time := time.Now().Format(time.ANSIC)
	// responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	// conn.Write([]byte(responseStr))

	// close conn
	// conn.Close()
}

func decodeBuffer(buffer []byte) {
	print(buffer)
	// return game_message

}
