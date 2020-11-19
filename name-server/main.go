package name

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

	fmt.Println("Naming server running!!")
	namingInvoker := NewNamingInvoker(&proxy)
	go namingInvoker.Invoke()

	fmt.Scanln()
}
