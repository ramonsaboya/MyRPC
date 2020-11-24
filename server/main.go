package server

import (
	"fmt"

	"github.com/ramonsaboya/myrpc/commons"
)

func Main() {
	proxy := commons.ClientProxy{
		Host:     "localhost",
		Port:     6666,
		Protocol: commons.TCP,
		ID:       1,
		TypeName: "Calculator",
	}

	nameService, err := NewNamingProxy(&proxy)

	if err != nil {
		panic(err)
	}
	proxy.Port = 6665
	reg, err := nameService.Register(proxy)

	if err != nil || *reg == false {
		panic("Failed to Register Calculator Service")
	}

	fmt.Println("Calculator server running!!")
	calculatorInvoker := NewCalculatorInvoker(&proxy, 3)
	go calculatorInvoker.Invoke()

	fmt.Scanln()
}
