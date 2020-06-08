package server

import (
	"crypto/rand"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
)

// Command type to be string able
type Command int32

// Export command conde enums
const (
	ISALIVE            int32 = 0
	PUT                int32 = 1
	GET                int32 = 2
	REMOVE             int32 = 3
	SHUTDOWN           int32 = 4
	WIPEOUT            int32 = 5
	GETPID             int32 = 6
	GETMEMBERSHIPCOUNT int32 = 7
)

// Export response code enums
const ()

// GenMsgID returns a random byte slice of the specified length
func GenMsgID(len int) []byte {
	id := make([]byte, len)
	rand.Read(id)
	return id
}

// MakeIsAliveRequest returns a protobuf message
// Checks if the server is still alive
func MakeIsAliveRequest() *pb.Msg {
	kvRequest := &pb.KVRequest{}
	kvRequest.Command = ISALIVE
	payload, _ := proto.Marshal(kvRequest)
	return makeMessage(payload)
}

func makeMessage(payload []byte) *pb.Msg {
	msg := &pb.Msg{}
	msg.MessageID = GenMsgID(16)
	msg.Payload = payload
	msg.CheckSum = GetCheckSum(msg)
	return msg
}
