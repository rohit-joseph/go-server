package checksum

import (
	"hash/crc64"
)

// GetCheckSum returns the checksum value given the messageID and payload
func GetCheckSum(messageID []byte, payload []byte) uint64 {
	cat := append(messageID, payload...)
	return crc64.Checksum(cat, crc64.MakeTable(crc64.ECMA))
}

// VerifyCheckSum returns if the checksum given is the same as the calculated one
func VerifyCheckSum(checksum uint64, messageID []byte, payload []byte) bool {
	return checksum == GetCheckSum(messageID, payload)
}
