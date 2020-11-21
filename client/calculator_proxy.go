package client

import (
	"github.com/mitchellh/mapstructure"
	"github.com/ramonsaboya/myrpc/commons"
)

type CalculatorProxy struct {
	proxy commons.ClientProxy
}

func NewCalculatorProxy(proxyRef *commons.ClientProxy) (*CalculatorProxy, error) {
	return &CalculatorProxy{
		proxy: *proxyRef,
	}, nil
}

func (calculator *CalculatorProxy) EquationRoots(a, b, c int) (*commons.EquationRoots, error) {
	params := make([]interface{}, 3)
	params[0] = a
	params[1] = b
	params[2] = c
	req := Request{Operation: "EquationRoots", Params: params}

	requestor, err := NewRequestor(&calculator.proxy)
	if err != nil {
		return nil, err
	}
	res, err := requestor.Invoke(req)
	if err != nil {
		return nil, err
	}

	resMap := res.(map[string]interface{})
	roots := commons.EquationRoots{}

	mapstructure.Decode(resMap, &roots)

	return &roots, nil
}
