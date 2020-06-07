package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	hlp "github.com/rohit-joseph/go-server/helpers"
	pb "github.com/rohit-joseph/go-server/proto"
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

	for {
		msg := &pb.Msg{}
		msg.MessageID = genMsgID(16)
		msg.Payload = genMsgID(25)
		msg.CheckSum = hlp.GetCheckSum(msg.MessageID, msg.Payload)

		out, err := proto.Marshal(msg)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(out)
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

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
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}

func genMsgID(len int) []byte {
	id := make([]byte, len)
	rand.Read(id)
	fmt.Println(id)
	return id
}
