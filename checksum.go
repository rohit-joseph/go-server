package server

import (
	"hash/crc32"

	pb "github.com/rohit-joseph/go-server/proto"
)

// GetCheckSum returns the checksum value given the messageID and payload
func GetCheckSum(msg *pb.Msg) uint32 {
	cat := append(msg.GetMessageID(), msg.GetPayload()...)
	return crc32.Checksum(cat, crc32.MakeTable(crc32.IEEE))
}

// VerifyCheckSum returns if the checksum given is the same as the calculated one
func VerifyCheckSum(msg *pb.Msg) bool {
	return msg.GetCheckSum() == GetCheckSum(msg)
}
