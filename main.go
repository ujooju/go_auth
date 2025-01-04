package main

import (
	"fmt"

	"github.com/ujooju/go_auth/server"
)

func main() {
	fmt.Println("Hello")
	server.StartServer("8081")
}
