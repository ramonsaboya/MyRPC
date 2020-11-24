package server

import (
	"errors"
	"fmt"
	"strconv"
)

type LifecycleManager struct {
	calculators map[string]Calculator
}

func NewLifecycleManager(poolSize int) *LifecycleManager {
	calculators := make(map[string]Calculator)
	for i := 0; i < poolSize; i++ {
		var id = "calc-" + strconv.Itoa(i)
		fmt.Println(id)
		calculators[id] = Calculator{id: id, available: true}
	}
	return &LifecycleManager{
		calculators: calculators,
	}
}

func (lm *LifecycleManager) Get() (Calculator, error) {
	for _, v := range lm.calculators {
		if v.available {
			v.available = false
			return v, nil
		}
	}
	return Calculator{}, errors.New("math: square root of negative number")
}

func (lm *LifecycleManager) Release(id string) {
	calculator := lm.calculators[id]
	calculator.available = true
	return
}
