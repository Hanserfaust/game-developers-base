package core

import "log"

type DispatcherMessage struct {
	// SourceID is for now the Username
	SourceID string
	Type     int
	Data     []byte
}

var ToClientChannels = make(map[string]chan DispatcherMessage)
var FromClientChannels = make(map[string]chan DispatcherMessage)

func GetUsernames() []string {
	return getAllKeysFromMap(ToClientChannels)
}

func getAllKeysFromMap(theMap map[string]chan DispatcherMessage) []string {
	keys := make([]string, 0, len(theMap))
	for k := range theMap {
		keys = append(keys, k)
	}
	return keys
}

func RegisterToClientChannel(username string, toClient chan DispatcherMessage) {
	ToClientChannels[username] = toClient
}

func RegisterFromClientChannel(username string, fromClient chan DispatcherMessage) {
	FromClientChannels[username] = fromClient

	go fromClientHandler(fromClient)
}

func toClientDispatcher(message DispatcherMessage) {
	// Look up the channel in the registry, and then send message
	toClientChannel := ToClientChannels[message.SourceID]
	toClientChannel <- message
}

func fromClientHandler(in chan DispatcherMessage) {
	for {
		dm := <-in
		log.Println("Dispatcher got message of type", dm.Type, "from:", dm.SourceID)

		// Route Message into Core etc.
	}
}
