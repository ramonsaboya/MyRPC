package main

import (
	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/commons"
	"github.com/ramonsaboya/myrpc/nameserver"
	"github.com/ramonsaboya/myrpc/server"
)

func main() {
	println(client.Which())
	println(commons.Which())
	println(nameserver.Which())
	println(server.Which())
}
