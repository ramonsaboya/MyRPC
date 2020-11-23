package client

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/ramonsaboya/myrpc/commons"
)

var iterations = 10000

func Main(protocol commons.Protocol, benchmark bool, wg sync.WaitGroup) {
	proxy := commons.ClientProxy{
		Host:     "localhost",
		Port:     6666,
		Protocol: protocol,
		ID:       1,
		TypeName: "Calculator",
	}

	nameService, err := NewNamingProxy(&proxy)

	if err != nil {
		panic(err)
	}

	calculatorProxy, err := nameService.Lookup("Calculator")

	if err != nil {
		panic(err)
	}

	calculator, err := NewCalculatorProxy(calculatorProxy)

	if err != nil {
		panic(err)
	}

	if benchmark {
		fmt.Println("benchmark")
		var sum int64 = 0
		iterationTime := make([]int64, iterations)
		for i := 0; i < iterations; i++ {
			startTime := time.Now()
			root, err := calculator.EquationRoots(2, 4, -6)
			totalTime := time.Now().Sub(startTime).Microseconds()
			fmt.Println(root)
			fmt.Println(totalTime)
			sum += totalTime
			iterationTime[i] = totalTime
			if err != nil {
				panic(err)
			}
		}
		var variation float64 = 0
		mean := float64(sum) / float64(iterations)
		for _, time := range iterationTime {
			diff := float64(time) - mean
			variation += diff * diff
		}
		variation /= float64(iterations)
		sd := math.Sqrt(variation)
		fmt.Println(mean)
		fmt.Println(sd)
		wg.Done()
	} else {
		for i := 0; i < iterations; i++ {
			_, err := calculator.EquationRoots(2, 4, -6)
			if err != nil {
				panic(err)
			}
		}
	}
	// fmt.Printf("roots are %v\n", roots.Roots)
}
