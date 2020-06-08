package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
	hlp "github.com/rohit-joseph/go-server/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	msg := makeIsAliveRequest()
	out, err := proto.Marshal(msg)
	_, err = c.Write(out)

	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	kvResponse := &pb.KVResponse{}
	err = proto.Unmarshal(buffer[0:n], kvResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(kvResponse.GetErrCode())
}

func makeIsAliveRequest() *pb.Msg {
	kvRequest := &pb.KVRequest{}
	kvRequest.Command = pb.KVRequest_CommandType_value["ISALIVE"]
	payload, _ := proto.Marshal(kvRequest)
	return makeMessage(payload)
}

func makeMessage(payload []byte) *pb.Msg {
	msg := &pb.Msg{}
	msg.MessageID = genMsgID(16)
	msg.Payload = payload
	msg.CheckSum = hlp.GetCheckSum(msg)
	return msg
}

func genMsgID(len int) []byte {
	id := make([]byte, len)
	rand.Read(id)
	return id
}
