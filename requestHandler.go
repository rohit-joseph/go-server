package server

import (
	"os"

	pb "github.com/rohit-joseph/go-server/proto"
)

// HandleRequest handles request
func HandleRequest(kvRequest *pb.KVRequest) *pb.KVResponse {
	kvResponse := &pb.KVResponse{}
	switch command := kvRequest.GetCommand(); command {
	case ISALIVE:
		kvResponse.ErrCode = SUCCESS
	case PUT:
		valVerPair := ValVerPair{
			value:   kvRequest.GetValue(),
			version: kvRequest.GetVersion(),
		}
		PutKVP(kvRequest.GetKey(), valVerPair)
		kvResponse.ErrCode = SUCCESS
	case GET:
		valVerPair := GetKVP(kvRequest.GetKey())
		if valVerPair.value == nil {
			kvResponse.ErrCode = NONEXISTENTKEY
		} else {
			kvResponse.ErrCode = SUCCESS
			kvResponse.Value = valVerPair.value
			kvResponse.Version = valVerPair.version
		}
	case SHUTDOWN:
		os.Exit(3)
	default:
		kvResponse.ErrCode = UNRECOGNIZEDCOMMAND
	}
	return kvResponse
}
