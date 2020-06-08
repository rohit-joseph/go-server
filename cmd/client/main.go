package main

import (
	"fmt"
	"os"

	srv "github.com/rohit-joseph/go-server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	srv.TestClient(CONNECT)
}
