package server

import (
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/rohit-joseph/go-server/proto"
)

func init() {
	// CreateLogger()
	CreateCache()
}

// Server running to respond to requests
func Server(PORT string) {
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		log.Println(err)
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Println(err)
	}

	defer connection.Close()
	buffer := make([]byte, 2048)
	rand.Seed(time.Now().Unix())
	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		var data []byte
		msg := &pb.Msg{}
		if err := proto.Unmarshal(buffer[0:n], msg); err != nil {
			log.Println(err)
			continue
		}

		if VerifyCheckSum(msg) == false {
			continue
		}

		if x, found := GetCache(msg.GetMessageID()); found {
			data = *x
		} else {
			kvRequest := &pb.KVRequest{}
			if err := proto.Unmarshal(msg.GetPayload(), kvRequest); err != nil {
				log.Println(err)
				continue
			}

			payload, err := proto.Marshal(HandleRequest(kvRequest))
			if err != nil {
				log.Println(err)
				continue
			}

			msg.Payload = payload
			msg.CheckSum = GetCheckSum(msg)

			data, err = proto.Marshal(msg)
			PutCache(msg.GetMessageID(), data)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

// func whatever(buffer []byte, n int, addr *net.UDPAddr) {
// 	var data []byte

// 	msg := &pb.Msg{}
// 	if err := proto.Unmarshal(buffer[0:n], msg); err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	if VerifyCheckSum(msg) == false {
// 		return
// 	}

// 	if x, found := GetCache(msg.GetMessageID()); found {
// 		data = *x
// 	} else {
// 		kvRequest := &pb.KVRequest{}
// 		if err := proto.Unmarshal(msg.GetPayload(), kvRequest); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		payload, err := proto.Marshal(HandleRequest(kvRequest))
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		msg.Payload = payload
// 		msg.CheckSum = GetCheckSum(msg)

// 		data, err = proto.Marshal(msg)
// 		PutCache(msg.GetMessageID(), data)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	s, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(rand.Intn(30000-20000+1)+20000))
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	connection, err := net.ListenUDP("udp4", s)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer connection.Close()

// 	_, err = connection.WriteToUDP(data, addr)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
