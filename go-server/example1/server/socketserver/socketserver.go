package socketserver

import (
	"bitknife.se/control"
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

	STATIC_PASSWORD = "#S33CR3T!"
)

func handleConnection(conn net.Conn, tDC chan control.DispatcherMessage, fDC chan control.DispatcherMessage) {
	/**
	Handles a connection from a Client.

	the receiver will move incoming data to the next place

	the sender will accept packages from other parts of the backend
	and send it to the client

	*/

	playerLogin := authenticateClient(conn)
	if playerLogin == nil {
		conn.Close()
		return
	}
	// Ok, Client is authenticated start routine to read packages from client

	// Register the needed channels
	setupChannels(playerLogin)

	// Main packet receiver
	go receivePacketsRoutine(conn, playerLogin, tDC)

	// Main packet sender
	go sendPackagesRoutine(conn)

}

func setupChannels(playerLogin *PlayerLogin) {

}

func receivePacketsRoutine(conn net.Conn, playerLogin *PlayerLogin, tDC chan control.DispatcherMessage) {
	for {
		messageType, messageData := receivePackageFromConnection(conn)

		if messageType == 0 {
			// Communication error?
			log.Println("ERROR from receivePackageFromConnection()!")

			// TODO: Improve
			conn.Close()
			return
		}

		// Ok got a valid message, pass that to the dispatcher
		dm := control.DispatcherMessage{SourceID: playerLogin.Username, Type: messageType, Data: messageData}
		tDC <- dm
	}
}

func authenticateClient(conn net.Conn) *PlayerLogin {
	/**
	Hard-coded login process. Will assume the correct order of packets
	sent from the client. Will authorize the client etc.

	Just assuming the PlayerLogin is the first
	package being sent is simple enough: If not, just disconnect.
	*/
	messageType, message := receivePackageFromConnection(conn)
	if messageType == int(MType_PLAYER_LOGIN) {
		playerLogin := bytesToPlayerLogin(messageType, message)

		// TODO: Check vs player-registry etc.
		if playerLogin.Password == STATIC_PASSWORD {
			log.Println("Username:", playerLogin.Username, "authenticatd.")
			return playerLogin
		} else {
			log.Println("ACCESS DENIED FOR Username:", playerLogin.Username, ": Invalid password.")
		}
	} else {
		log.Println("ACCESS DENIED: Invalid message type", messageType, "when authenticating.")
	}
	return nil
}

func receivePackageFromConnection(conn net.Conn) (int, []byte) {
	/**
	Waits for the header and returns the type and []byte representing the package.
	*/
	// printReceivedBuffer(messageData, messageType)

	// Allocate header
	header := make([]byte, 2)

	// First read the two byte header
	_, err := io.ReadAtLeast(conn, header, 2)

	if err != nil {
		// Broken connection, client ugly shutdown etc.
		log.Print("Error reading from:", conn.RemoteAddr(), "reason was: ", err)
		log.Print("Closing!", conn)
		return 0, nil
	}

	messageSize := header[0]
	messageType := int(header[1])

	// Allocate for packet
	messageData := make([]byte, messageSize)

	// And read the packet
	_, err = io.ReadFull(conn, messageData)

	return messageType, messageData
}

func bytesToPlayerLogin(messageType int, messageData []byte) *PlayerLogin {
	/**
	This one is a bit odd. Should really not have "gameMessages" on this layer
	but this simplifies things a bit for now.

	A solution would be to generalize the message concept into game messages
	and all other kinds of messages supporting (ie login, logout, ping etc.)
	*/
	gameMessage := packetToGameMessage(messageData, messageType)
	playerLogin := gameMessage.(PlayerLogin)
	return &playerLogin
}

func sendPackagesRoutine(conn net.Conn) {

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

func Start(tDC chan control.DispatcherMessage, fDC chan control.DispatcherMessage) {

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		// Note: call to Fatal will do os.Exit(1).
		log.Fatal(err)
	}
	fmt.Println("Server up on", HOST, ":", PORT)

	// close listener
	defer listen.Close()
	for {
		log.Println("Waiting for clients to connect...")
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Failed to Accept():", err)
		}
		log.Println("Client connected from", conn.RemoteAddr())

		go handleConnection(conn, tDC, fDC)
	}
}
