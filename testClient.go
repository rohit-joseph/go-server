package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
)

var c *net.UDPConn
var timeout = 100

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
	shutdown()
}

func testIsAliveRequest() {
	testName := "ISALIVE TEST: "
	msg := MakeIsAliveRequest()
	msg = requestReply(msg)

	if msg == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	s := getErrCode(msg)
	fmt.Println(testName + stringErrCode(s))
}

func testPutGetRequest() {
	testName := "PUT and GET TEST: "
	key := GenRandomSlice(16)
	value := GenRandomSlice(100)
	var version int32 = 0

	msg := MakePutRequest(key, value, version)
	msg = requestReply(msg)

	if msg == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(msg); s != SUCCESS {
		fmt.Println(testName + stringErrCode(s))
		return
	}

	msg = MakeGetRequest(key)
	msg = requestReply(msg)
	if msg == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(msg); s == SUCCESS {
		kvResponse := getdPayload(msg)

		if string(kvResponse.GetValue()) != string(value) || kvResponse.GetVersion() != version {
			fmt.Println(testName + "PUT and GET values did not match")
		} else {
			fmt.Println(testName + stringErrCode(s))
		}
	} else {
		fmt.Println(testName + stringErrCode(s))
	}

}

func shutdown() {
	testName := "SHUTDOWN TEST: "
	msg := MakeShutDownRequest()
	writeToConnection(msg)

	msg = MakeIsAliveRequest()
	msg = requestReply(msg)

	if msg == nil {
		fmt.Println(testName + "SUCCESS")
		return
	}
}

func requestReply(msg *pb.Msg) *pb.Msg {
	i := 1
	for i <= 4 {
		writeToConnection(msg)
		delay := time.Duration(timeout * i * 2)
		c.SetReadDeadline(time.Now().Add(time.Millisecond * delay))
		ret := readFromConnection()
		if ret != nil {
			return ret
		}
		i++
	}
	return nil
}

func getErrCode(msg *pb.Msg) uint32 {
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	return kvResponse.GetErrCode()
}

func getdPayload(msg *pb.Msg) *pb.KVResponse {
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	return kvResponse
}

func writeToConnection(msg *pb.Msg) {
	out, _ := proto.Marshal(msg)
	_, _ = c.Write(out)
}

func readFromConnection() *pb.Msg {
	buffer := make([]byte, 1024)
	n, _, err := c.ReadFromUDP(buffer)
	if err != nil {
		log.Println(err)
		return nil
	}

	msg := &pb.Msg{}
	err = proto.Unmarshal(buffer[0:n], msg)
	log.Println(err)
	return msg
}
