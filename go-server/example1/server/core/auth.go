package core

import (
	"log"
)

const (
	STATIC_PASSWORD = "#S33CR3T!"
)

func AuthenticateClient(messageType int, message []byte) *PlayerLogin {
	/**
	Hard-coded login process. Will assume the correct order of packets
	sent from the client. Will authorize the client etc.

	Just assuming the PlayerLogin is the first
	package being sent is simple enough: If not, just disconnect.
	*/

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

func bytesToPlayerLogin(messageType int, messageData []byte) *PlayerLogin {
	/**
	Works for now.
	*/
	gameMessage := PacketToGameMessage(messageData, messageType)
	playerLogin := gameMessage.(PlayerLogin)
	return &playerLogin
}
