package socketserver

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestProtocolBuffers(t *testing.T) {
	playerLogin := PlayerLogin{
		Username: "Tester",
		Password: "s33cret!",
	}
	out, err := proto.Marshal(&playerLogin)

	if err != nil {
		log.Fatalln("Failed to Marshal PlayerLogin object:", err)
	}
	encodedStr := hex.EncodeToString(out)
	fmt.Println("Encoded to: ", encodedStr)

	playerLoginCopy := PlayerLogin{}
	err = proto.Unmarshal(out, &playerLoginCopy)

	fmt.Println("Username: ", playerLoginCopy.Username)
	fmt.Println("Password: ", playerLoginCopy.Password)
}