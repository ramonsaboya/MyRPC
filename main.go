package main

import (
	"os"

	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/server"
)

func main() {
	service := os.Args[1]

	switch service {
	case "client":
		client.Main()
	case "server":
		server.Main()
	}
}
