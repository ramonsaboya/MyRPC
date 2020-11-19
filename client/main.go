package client

import (
	"fmt"

	nameserver "github.com/ramonsaboya/myrpc/name-server"
)

func Main() {
	namingService := nameserver.NameServer{}

	calculatorProxy, _ := namingService.Lookup("Calculator")

	calculator, err := NewCalculatorProxy(&calculatorProxy)
	if err != nil {
		panic(err)
	}

	roots, err := calculator.EquationRoots(2, 4, -6)
	if err != nil {
		panic(err)
	}
	fmt.Printf("roots are %v\n", roots.Roots)
}
