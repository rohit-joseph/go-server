package server

import (
	"hash/crc64"

	pb "github.com/rohit-joseph/go-server/proto"
)

// GetCheckSum returns the checksum value given the messageID and payload
func GetCheckSum(msg *pb.Msg) uint64 {
	cat := append(msg.GetMessageID(), msg.GetPayload()...)
	return crc64.Checksum(cat, crc64.MakeTable(crc64.ECMA))
}

// VerifyCheckSum returns if the checksum given is the same as the calculated one
func VerifyCheckSum(msg *pb.Msg) bool {
	return msg.GetCheckSum() == GetCheckSum(msg)
}
