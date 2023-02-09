package main

import (
	"bitknife.se/gateserver"
)

func main() {
	// Socket server
	go gateserver.Start()

	// TODO: start the game server
	// go gameserver.Start()

	// TODO: start the meta-services integration stuff
	// go metaserver.Start()
}
