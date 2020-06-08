package server

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
)

var c *net.UDPConn

// TestClient runs the test client
func TestClient(CONNECT string) {
	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err = net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	testIsAliveRequest()
}

func testIsAliveRequest() {
	msg := MakeIsAliveRequest()
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

	if kvResponse.GetErrCode() != 0 {
		log.Println(kvResponse.GetErrCode())
	}
}
