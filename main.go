package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ramonsaboya/myrpc/commons"

	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/name-server"
	"github.com/ramonsaboya/myrpc/server"
)

func main() {
	service := os.Args[1]
	protocol := os.Args[2]
	var ClientAmounts = []int{1, 2, 5, 10}
	var wg sync.WaitGroup
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
			wg.Add(1)
			go client.Main(_protocol, true, &wg)
			for i := 0; i < clientAmount-1; i++ {
				go client.Main(_protocol, false, &wg)
			}
			wg.Wait()
			fmt.Println("###############")
		}
	case "server":
		server.Main(_protocol)
	case "name":
		name.Main(_protocol)
	}
}
