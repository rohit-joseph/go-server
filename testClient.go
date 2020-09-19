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
var sendAddr *net.UDPAddr
var timeout = 100
var err error

// TestClient runs the test client
func TestClient(CONNECT string) {
	sendAddr, err = net.ResolveUDPAddr("udp4", CONNECT)
	testClientAddr, err := net.ResolveUDPAddr("udp4", ":55555")
	c, err = net.ListenUDP("udp4", testClientAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The UDP server is %s\n", CONNECT)
	defer c.Close()

	// testIsAliveRequest()
	// testPutGetRequest()
	// testWipeout()
	// testGetPid()
	// testGetMembershipCount()
	// testAtMostOnceSemantics()
	performanceTest(1, 60)
	performanceTest(2, 60)
	performanceTest(4, 60)
	performanceTest(8, 60)
	// shutdown()
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
		kvResponse := getPayload(msg)

		if string(kvResponse.GetValue()) != string(value) || kvResponse.GetVersion() != version {
			fmt.Println(testName + "PUT and GET values did not match")
		} else {
			fmt.Println(testName + stringErrCode(s))
		}
	} else {
		fmt.Println(testName + stringErrCode(s))
	}
}

func testWipeout() {
	testName := "WIPEOUT TEST: "
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

	msg = MakeWipeoutRequest()
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

	if s := getErrCode(msg); s == NONEXISTENTKEY {
		fmt.Println(testName + "SUCCESS")
	} else {
		fmt.Println(testName + "FAILED with err code: " + stringErrCode(s))
	}
}

func testGetPid() {
	testName := "GETPID TEST: "
	msg := MakeGetPID()
	msg = requestReply(msg)

	if msg == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(msg); s == SUCCESS {
		kvResponse := getPayload(msg)

		if kvResponse.GetPid() != 0 {
			fmt.Println(testName + stringErrCode(s))
		} else {
			fmt.Println(testName + "FAILED")
		}
	} else {
		fmt.Println(testName + stringErrCode(s))
	}
}

func testGetMembershipCount() {
	testName := "GETMEMBERSHIPCOUNT TEST: "
	msg := MakeGetMembershipCount()
	msg = requestReply(msg)

	if msg == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(msg); s == SUCCESS {
		kvResponse := getPayload(msg)

		if kvResponse.GetMembershipCount() == 1 {
			fmt.Println(testName + stringErrCode(s))
		} else {
			fmt.Println(testName + "FAILED")
		}
	} else {
		fmt.Println(testName + stringErrCode(s))
	}
}

func testAtMostOnceSemantics() {
	testName := "ATMOSTONCESEMANTICS TEST: "
	key := GenRandomSlice(16)
	value := GenRandomSlice(100)
	var version int32 = 0

	putMessage := MakePutRequest(key, value, version)
	putMessage = requestReply(putMessage)
	if putMessage == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(putMessage); s != SUCCESS {
		fmt.Println(testName + stringErrCode(s))
		return
	}

	removeMsg1 := MakeRemoveRequest(key)
	removeMsg1 = requestReply(removeMsg1)
	if removeMsg1 == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(removeMsg1); s != SUCCESS {
		fmt.Println(testName + stringErrCode(s))
		return
	}

	removeMsg2 := MakeRemoveRequest(key)
	removeMsg2 = requestReply(removeMsg2)
	if removeMsg2 == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}

	if s := getErrCode(removeMsg2); s != INVALIDKEY {
		fmt.Println(testName + "FAILED")
		return
	}

	removeMsg1 = requestReply(removeMsg1)
	if removeMsg1 == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}
	if s := getErrCode(removeMsg1); s != SUCCESS {
		fmt.Println(testName + stringErrCode(s))
		return
	}

	removeMsg2 = requestReply(removeMsg2)
	if removeMsg2 == nil {
		fmt.Println(testName + "TIMEOUT")
		return
	}
	if s := getErrCode(removeMsg2); s != INVALIDKEY {
		fmt.Println(testName + "FAILED. Got response: " + stringErrCode(s))
		return
	}

	fmt.Println(testName + "SUCCESS")
}

func performanceTest(numClients int, duration int) {
	totalRequest := 0
	totalSuccess := 0
	totalTimeout := 0
	totalFail := 0
	totalRequestAccumulator := make(chan int, numClients)
	totalSuccessAccumulator := make(chan int, numClients)
	totalTimeoutAccumulator := make(chan int, numClients)
	totalFailAccumulator := make(chan int, numClients)
	for i := 0; i < numClients; i++ {
		go requestRoutine(duration, totalRequestAccumulator, totalSuccessAccumulator, totalTimeoutAccumulator, totalFailAccumulator)
	}

	for j := 0; j < numClients; j++ {
		totalRequest += <-totalRequestAccumulator
		totalSuccess += <-totalSuccessAccumulator
		totalTimeout += <-totalTimeoutAccumulator
		totalFail += <-totalFailAccumulator
	}

	fmt.Printf("Number of requests: %d\n", totalRequest)
	fmt.Printf("Number of success: %d\n", totalSuccess)
	fmt.Printf("Number of timeouts: %d\n", totalTimeout)
	fmt.Printf("Number of fails: %d\n", totalFail)
	fmt.Printf("Number of requests/second: %d\n", (totalRequest / duration))
	fmt.Printf("Number of successful requests/second: %d\n", (totalSuccess / duration))
}

func requestRoutine(duration int, requestAcc chan int, successAcc chan int, timeoutAcc chan int, failAcc chan int) {
	startTime := time.Now()
	count := 0
	success := 0
	timeout := 0
	fail := 0
	for time.Now().Before(startTime.Add(time.Second * time.Duration(duration))) {
		count++
		key := GenRandomSlice(16)
		value := GenRandomSlice(100)
		var version int32 = 0

		msg := MakePutRequest(key, value, version)
		msg = requestReply(msg)

		if msg == nil {
			timeout++
			continue
		}

		if s := getErrCode(msg); s != SUCCESS {
			fail++
			continue
		}

		msg = MakeGetRequest(key)
		msg = requestReply(msg)
		if msg == nil {
			timeout++
			continue
		}

		if s := getErrCode(msg); s == SUCCESS {
			kvResponse := getPayload(msg)

			if string(kvResponse.GetValue()) != string(value) || kvResponse.GetVersion() != version {
				fail++
			}
		}

		msg = MakeRemoveRequest(key)
		msg = requestReply(msg)
		if msg == nil {
			timeout++
			continue
		}

		if s := getErrCode(msg); s != SUCCESS {
			fail++
			continue
		}
		success++
	}
	requestAcc <- count
	successAcc <- success
	timeoutAcc <- timeout
	failAcc <- fail
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
			if verifyMessage(msg, ret) {
				return ret
			}
		}
		i++
	}
	return nil
}

func getErrCode(msg *pb.Msg) int32 {
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	return kvResponse.GetErrCode()
}

func getPayload(msg *pb.Msg) *pb.KVResponse {
	kvResponse := &pb.KVResponse{}
	_ = proto.Unmarshal(msg.GetPayload(), kvResponse)

	return kvResponse
}

func writeToConnection(msg *pb.Msg) {
	out, _ := proto.Marshal(msg)
	_, _ = c.WriteToUDP(out, sendAddr)
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
	if err != nil {
		log.Println(err)
		return nil
	}
	return msg
}

func verifyMessage(org *pb.Msg, ret *pb.Msg) bool {
	if (string(org.MessageID) == string(ret.MessageID)) && VerifyCheckSum(ret) {
		return true
	}
	return false
}
