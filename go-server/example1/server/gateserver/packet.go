package gateserver

import (
	"fmt"
	"google.golang.org/protobuf/proto"
)

func packetToGameMessage(buffer []byte, mType int) interface{} {

	fmt.Println("Got Game Message of type: ", mType)

	switch mType {

	case int(MType_PLAYER_LOGIN):
		{
			message := PlayerLogin{}
			proto.Unmarshal(buffer, &message)
			return message
		}

	case int(MType_MOUSE_EVENT):
		{
			message := MouseEvent{}
			proto.Unmarshal(buffer, &message)
			return message
		}
	}
	return nil
}
