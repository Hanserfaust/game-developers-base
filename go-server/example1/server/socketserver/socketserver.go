package socketserver

import (
	"bitknife.se/core"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "7777"
	TYPE = "tcp"

	STATIC_PASSWORD = "#S33CR3T!"
)

func handleConnection(conn net.Conn) {
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

	/**
	Client is authenticated, now we need to connect the client
	to the game. This is done using Channels that connects to
	the Dispatcher (middle layer), which then in turn connects to
	the game engine (upper layer).

	This separates the socket layer from the game layer completely.
	*/

	// Create and register the needed channels on the dispatcher
	fromClient, toClient := makeAndRegisterChannels(playerLogin)

	// Main packet receiver
	go receivePacketsRoutine(conn, playerLogin, fromClient)

	// Main packet sender
	go sendPackagesRoutine(conn, toClient)

}

func makeAndRegisterChannels(playerLogin *core.PlayerLogin) (chan core.DispatcherMessage, chan core.DispatcherMessage) {
	fromClient := make(chan core.DispatcherMessage)
	toClient := make(chan core.DispatcherMessage)

	// And register channels on the Dispatcher
	core.RegisterToClientChannel(playerLogin.Username, toClient)
	core.RegisterFromClientChannel(playerLogin.Username, fromClient)

	return fromClient, toClient
}

func receivePacketsRoutine(conn net.Conn, playerLogin *core.PlayerLogin, fromClient chan core.DispatcherMessage) {
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
		dm := core.DispatcherMessage{SourceID: playerLogin.Username, Type: messageType, Data: messageData}
		fromClient <- dm
	}
}

func authenticateClient(conn net.Conn) *core.PlayerLogin {
	/**
	Hard-coded login process. Will assume the correct order of packets
	sent from the client. Will authorize the client etc.

	Just assuming the PlayerLogin is the first
	package being sent is simple enough: If not, just disconnect.
	*/
	messageType, message := receivePackageFromConnection(conn)
	if messageType == int(core.MType_PLAYER_LOGIN) {
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

func bytesToPlayerLogin(messageType int, messageData []byte) *core.PlayerLogin {
	/**
	This one is a bit odd. Should probably not have "gameMessages" on this layer
	but this simplifies things a bit for now.

	A solution would be to generalize the message concept into game messages
	and all other kinds of messages supporting (ie login, logout, ping etc.)
	*/
	gameMessage := core.PacketToGameMessage(messageData, messageType)
	playerLogin := gameMessage.(core.PlayerLogin)
	return &playerLogin
}

func sendPackagesRoutine(conn net.Conn, toClient chan core.DispatcherMessage) {
	for {
		dm := <-toClient

		packet := dm.Data

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
		log.Println("Waiting for clients to connect...")
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Failed to Accept():", err)
		}
		log.Println("Client connected from", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
