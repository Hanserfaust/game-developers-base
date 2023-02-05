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

func handleConnection(conn net.Conn) {
	/**
	Outer routine called after the Socket Accept.

	This one must register the connection somehow
	on a central place since other routines will need
	to send messages to the Player.

	This routine:
		- expect a PlayerLogin.
		- check the credentials.
		- send back PlayerLoginSuccess or PlayerLoginFailed
		- if Failed: clean up and close the connection.
		- if Success: register the player on the World.
	*/

	for {
		// Allocate header
		header := make([]byte, 2)

		// First read the two byte header
		_, err := io.ReadAtLeast(conn, header, 2)

		if err != nil {
			log.Fatal(err)
		}

		messageSize := header[0]
		messageType := int(header[1])

		// Allocate for packet
		messageData := make([]byte, messageSize)

		// And read the packet
		_, err = io.ReadFull(conn, messageData)

		// printReceivedBuffer(messageData, messageType)

		// Unmarshal binary data into Protocol Buffer gamepacket
		gameMessage := packetToGameMessage(messageData, messageType)

		if messageType == int(MType_PLAYER_LOGIN) {
			playerLogin := gameMessage.(PlayerLogin)

			// TODO: Check credentials

			fmt.Println("Logged in: ", playerLogin.Username)

		} else {
			fmt.Println("!!! EVENT FROM NON LOGGED IN CONNECTION !!!")
			// conn.Close()
		}
		// fmt.Println(game_message.name)
	}

	// write data to response
	// time := time.Now().Format(time.ANSIC)
	// responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	// conn.Write([]byte(responseStr))

	// close conn
	// conn.Close()
}

func printReceivedBuffer(buffer []byte, messageType int) {
	fmt.Printf("Got message of type 0x%x:", messageType)

	encodedStr := hex.EncodeToString(buffer)
	fmt.Printf("%s\n", encodedStr)
}

func toGameMessage(buffer []byte, messageType byte) {
}

func Start() {

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
			// os.Exit(1)
		}

		go handleConnection(conn)
	}
}
