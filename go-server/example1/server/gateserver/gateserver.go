package gateserver

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "7777"
	TYPE = "tcp"
)

func handleConnection(conn net.Conn) {
	/**
	Handles a connection from a Client.

	the receiver will move incoming data to the next place

	the sender will accept packages from other parts of the backend
	and send it to the client

	*/

	// Main packet receiver
	go receivePackagesFromConnection(conn)

	// Main packet sender
	go sendPackagesToConnection(conn)
}

func receivePackagesFromConnection(conn net.Conn) {
	/**
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

		// TODO: Move message to right input queue/goroutine based on type?

		// For now
		if messageType == int(MType_PLAYER_LOGIN) {
			playerLogin := gameMessage.(PlayerLogin)

			// TODO: Check credentials etc.

			fmt.Println("Logged in: ", playerLogin.Username)

		} else {
			fmt.Println("Dropping event of type", messageType)
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

func sendPackagesToConnection(conn net.Conn) {

	for {
		time.Sleep(100 * time.Microsecond)

		fmt.Println("Sending PING")

		packet := buildPingPacket()

		_, err := conn.Write(packet)
		if err != nil {
			return
		}
	}
}

func Start() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Server up on", HOST, ":", PORT)

	// close listener
	// defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			// os.Exit(1)
		}

		go handleConnection(conn)
	}
}
