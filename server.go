package server

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
)

// Server running to respond to requests
func Server(PORT string) {
	s, err := net.ResolveUDPAddr("udp4", PORT)
	printErr(err)

	connection, err := net.ListenUDP("udp4", s)

	printErr(err)

	defer connection.Close()
	buffer := make([]byte, 2048)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)

		msg := &pb.Msg{}
		if err := proto.Unmarshal(buffer[0:n], msg); err != nil {
			fmt.Println(err)
			continue
		}

		if VerifyCheckSum(msg) == false {
			continue
		}

		kvRequest := &pb.KVRequest{}
		if err = proto.Unmarshal(msg.GetPayload(), kvRequest); err != nil {
			fmt.Println(err)
			continue
		}

		payload, err := proto.Marshal(HandleRequest(kvRequest))
		printErr(err)

		msg.Payload = payload
		msg.CheckSum = GetCheckSum(msg)

		data, err := proto.Marshal(msg)
		printErr(err)

		_, err = connection.WriteToUDP(data, addr)
		printErr(err)
	}
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
