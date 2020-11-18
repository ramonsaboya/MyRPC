package client

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

	calculator, err := NewCalculatorProxy(&proxy)
	if err != nil {
		panic(err)
	}

	roots, err := calculator.EquationRoots(2, 4, -6)
	if err != nil {
		panic(err)
	}
	fmt.Printf("roots are %v\n", roots.Roots)
}
