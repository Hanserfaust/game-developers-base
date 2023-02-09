package gateserver

import (
	"fmt"
	"io"
	"log"
	"net"
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
			// Broken connection, client ugly shutdown etc.
			log.Print("Error reading from client connection:", err)
			log.Print("Closing!", conn)
			return
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

			log.Println("Logged in:", playerLogin.Username)

		} else {
			// fmt.Println("Dropping event of type", messageType)
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
		time.Sleep(1000 * time.Millisecond)

		packet := buildPingPacket()

		_, err := conn.Write(packet)
		if err != nil {
			log.Println("Error writing packet: ")
			return
		}
	}
}

func Start() {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		// Note: call to Fatal will do os.Exit(1).
		log.Fatal(err)
	}
	fmt.Println("Server up on", HOST, ":", PORT)

	// close listener
	defer listen.Close()
	for {
		log.Println("Listening for client.")
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Failed to Accept():", err)
		}
		log.Println("Client connected from", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
