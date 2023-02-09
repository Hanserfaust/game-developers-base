package main

import (
	"bitknife.se/core"
	"bitknife.se/socketserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func startServer() {

	// Fancy console for the future!
	// go StartConsole()

	/**
	Handles the TCP connections, moving messages through
	the channels.
	*/
	go socketserver.Start()

	/**
	This is the main game logic core.
	*/
	go core.Start()

	// TODO:
	// Start REST APi etc.
	// Start meta-services integration stuff (logging, metrics etc.)
}

func stopServer() {
	log.Println("Stopping server.")

	// TODO: call goroutines to shutdown properly
}

func mainLoop() {
	exitSignal := make(chan os.Signal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	stopServer()
}

func main() {

	// Spawns everything we need
	startServer()

	// Waits for SIGINT and SIGTERM to perform shutdown
	mainLoop()

}
