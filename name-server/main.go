package name

import (
	"fmt"

	"github.com/ramonsaboya/myrpc/commons"
)

func Main(protocol commons.Protocol) {
	proxy := commons.ClientProxy{
		Host:     "localhost",
		Port:     6666,
		Protocol: protocol,
		ID:       1,
		TypeName: "Calculator",
	}

	fmt.Println("Naming server running!!")
	namingInvoker := NewNamingInvoker(&proxy)
	go namingInvoker.Invoke()

	fmt.Scanln()
}
