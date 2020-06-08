package main

import (
	"fmt"
	"math/rand"
	"os"

	srv "github.com/rohit-joseph/go-server"
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

	srv.Server(PORT)
}
