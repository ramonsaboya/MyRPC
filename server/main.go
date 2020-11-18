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

	fmt.Println("Calculator server running!!")
	calculatorInvoker := NewCalculatorInvoker(&proxy)
	go calculatorInvoker.Invoke()

	fmt.Scanln()
}
