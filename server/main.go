package server

import (
	"fmt"

	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/commons"
	nameserver "github.com/ramonsaboya/myrpc/name-server"
)

func Main() {
	proxy := commons.ClientProxy{
		Host:     "localhost",
		Port:     6666,
		Protocol: commons.TCP,
		ID:       1,
		TypeName: "Calculator",
	}

	nameService := nameserver.NameServer{}

	client.NewCalculatorProxy(&proxy)

	nameService.Register("Calculator", proxy)

	fmt.Println("Calculator server running!!")
	calculatorInvoker := NewCalculatorInvoker(&proxy)
	go calculatorInvoker.Invoke()

	fmt.Scanln()
}
