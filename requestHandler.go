package server

import (
	pb "github.com/rohit-joseph/go-server/proto"
)

// HandleRequest handles request
func HandleRequest(kvRequest *pb.KVRequest) *pb.KVResponse {
	kvResponse := &pb.KVResponse{}
	switch command := kvRequest.GetCommand(); command {
	case ISALIVE:
		kvResponse.ErrCode = 0
	default:
		kvResponse.ErrCode = 5
	}

	kvResponse.ErrCode = 0
	return kvResponse
}
