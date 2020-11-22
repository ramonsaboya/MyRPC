package main

import (
	"fmt"
	"os"

	"github.com/ramonsaboya/myrpc/commons"

	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/name-server"
	"github.com/ramonsaboya/myrpc/server"
)

var ClientAmounts = []int{1, 2, 5, 10}

func main() {
	service := os.Args[1]
	protocol := os.Args[2]
	var _protocol commons.Protocol
	if protocol == "tcp" {
		_protocol = commons.TCP
	} else {
		_protocol = commons.UDP

	}
	switch service {
	case "client":
		for _, clientAmount := range ClientAmounts {
			fmt.Println("###############")
			fmt.Println(clientAmount)
			go client.Main(_protocol, true)
			for i := 0; i < clientAmount-1; i++ {
				go client.Main(_protocol, false)
			}
			fmt.Println("###############")
		}
		fmt.Scanln()
	case "server":
		server.Main(_protocol)
	case "name":
		name.Main(_protocol)
	}
}
