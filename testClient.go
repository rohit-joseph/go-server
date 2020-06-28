package server

import (
	"fmt"
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
	testPutGetRequest()
}

func testIsAliveRequest() {
	msg := MakeIsAliveRequest()

	writeToConnection(msg)

	msg = readFromConnection()

	if s := checkSuccess(msg); s == true {
		fmt.Println("Test: is alive success")
	} else {
		fmt.Println("Failed is alive")
	}
}

func testPutGetRequest() {
	key := GenRandomSlice(16)
	value := GenRandomSlice(100)
	var version int32 = 0

	msg := MakePutRequest(key, value, version)
	writeToConnection(msg)

	msg = readFromConnection()

	checkSuccess(msg)

	msg = MakeGetRequest(key)
	writeToConnection(msg)

	msg = readFromConnection()

	checkSuccess(msg)
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	if string(kvResponse.GetValue()) != string(value) || kvResponse.GetVersion() != version {
		fmt.Println("Failed put and get")
	} else {
		fmt.Println("Test: put and get success")
	}
}

func checkSuccess(msg *pb.Msg) bool {
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	return kvResponse.GetErrCode() == 0
}

func writeToConnection(msg *pb.Msg) {
	out, _ := proto.Marshal(msg)
	_, _ = c.Write(out)
}

func readFromConnection() *pb.Msg {
	buffer := make([]byte, 1024)
	n, _, _ := c.ReadFromUDP(buffer)

	msg := &pb.Msg{}
	_ = proto.Unmarshal(buffer[0:n], msg)
	return msg
}
