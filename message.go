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
const (
	SUCCESS                 uint32 = 0
	NONEXISTENTKEY          uint32 = 1
	OUTOFSPACE              uint32 = 2
	TEMPORARYSYSTEMOVERLOAD uint32 = 3
	INTERNALKVSTOREFAILURE  uint32 = 4
	UNRECOGNIZEDCOMMAND     uint32 = 5
	INVALIDKEY              uint32 = 6
	INVALIDVALUE            uint32 = 7
)

// GenRandomSlice returns a random byte slice of the specified length
func GenRandomSlice(len int) []byte {
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

// MakePutRequest returns a protobuf message
// Puts a key, value, and version into the kvStore
func MakePutRequest(key []byte, value []byte, version int32) *pb.Msg {
	kvRequest := &pb.KVRequest{}
	kvRequest.Command = PUT
	kvRequest.Key = key
	kvRequest.Value = value
	kvRequest.Version = version
	payload, _ := proto.Marshal(kvRequest)
	return makeMessage(payload)
}

// MakeGetRequest returns a protobuf message
// Gets a key from the kvStore
func MakeGetRequest(key []byte) *pb.Msg {
	kvRequest := &pb.KVRequest{}
	kvRequest.Command = GET
	kvRequest.Key = key
	payload, _ := proto.Marshal(kvRequest)
	return makeMessage(payload)
}

// MakeRemoveRequest returns a protobuf message
// Removes a key from the kvStore
func MakeRemoveRequest(key []byte) *pb.Msg {
	kvRequest := &pb.KVRequest{}
	kvRequest.Command = REMOVE
	kvRequest.Key = key
	payload, _ := proto.Marshal(kvRequest)
	return makeMessage(payload)
}

func makeMessage(payload []byte) *pb.Msg {
	msg := &pb.Msg{}
	msg.MessageID = GenRandomSlice(16)
	msg.Payload = payload
	msg.CheckSum = GetCheckSum(msg)
	return msg
}
