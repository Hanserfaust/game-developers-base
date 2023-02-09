package main

import (
	"bitknife.se/control"
	"bitknife.se/core"
	"bitknife.se/socketserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func startServer() {

	// Fancy console for the future
	// go StartConsole()
	tDC := make(chan control.DispatcherMessage)
	fDC := make(chan control.DispatcherMessage)

	go control.MainDispatcher(tDC, fDC)

	// TODO: Move into dispatcher module?

	go socketserver.Start(tDC, fDC)

	go core.Start()

	// Start REST APi etc.

	// Start meta-services integration stuff,
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
