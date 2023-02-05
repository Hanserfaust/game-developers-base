package gateserver

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
		Password: "foobar",
	}
	out, err := proto.Marshal(&playerLogin)

	if err != nil {
		log.Fatalln("Failed to Marshal PlayerLogin object:", err)
	}
	encodedStr := hex.EncodeToString(out)
	fmt.Printf("Encoded to: %s\n", encodedStr)
}
