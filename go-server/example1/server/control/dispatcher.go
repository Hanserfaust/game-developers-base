package control

import "log"

type DispatcherMessage struct {
	// SourceID is for now the Username
	SourceID string
	Type     int
	Data     []byte
}

func MainDispatcher(fromClients chan DispatcherMessage, toClients chan DispatcherMessage) {
	go fromClientsDispatcher(fromClients)
	go toClientsDispatcher(toClients)
}

func toClientsDispatcher(out chan DispatcherMessage) {
}

func fromClientsDispatcher(in chan DispatcherMessage) {
	for {
		dm := <-in
		log.Println("Dispatcher got message of type", dm.Type, "from:", dm.SourceID)
	}
}
