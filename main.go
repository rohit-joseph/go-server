package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
	hlp "github.com/rohit-joseph/go-server/server"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

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

		if hlp.VerifyCheckSum(msg) == false {
			continue
		}

		kvRequest := &pb.KVRequest{}
		if err = proto.Unmarshal(msg.GetPayload(), kvRequest); err != nil {
			fmt.Println(err)
			continue
		}

		payload, err := proto.Marshal(handleRequest(kvRequest))
		printErr(err)

		msg.Payload = payload
		msg.CheckSum = hlp.GetCheckSum(msg)

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

func handleRequest(kvRequest *pb.KVRequest) *pb.KVResponse {
	kvResponse := &pb.KVResponse{}
	switch command := getCommand(kvRequest); command {
	case "ISALIVE":
		kvResponse.ErrCode = 0
	default:
		kvResponse.ErrCode = 5
	}

	kvResponse.ErrCode = 0
	return kvResponse
}

func getCommand(kvRequest *pb.KVRequest) string {
	return pb.KVRequest_CommandType_name[kvRequest.GetCommand()]
}
