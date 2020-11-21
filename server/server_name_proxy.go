package server

import (
	"github.com/ramonsaboya/myrpc/client"
	"github.com/ramonsaboya/myrpc/commons"
)

type NameProxy struct {
	proxy commons.ClientProxy
}

func NewNamingProxy(proxyRef *commons.ClientProxy) (*NameProxy, error) {
	return &NameProxy{
		proxy: *proxyRef,
	}, nil
}

func (n *NameProxy) Register(proxy commons.ClientProxy) (*bool, error) {
	params := make([]interface{}, 2)
	params[0] = "Calculator"
	params[1] = proxy
	req := client.Request{
		Operation: "Register",
		Params:    params,
	}

	requestor, err := client.NewRequestor(&n.proxy)
	if err != nil {
		return nil, err
	}
	res, err := requestor.Invoke(req)

	if err != nil {
		return nil, err
	}

	reply := res.(bool)

	return &reply, nil
}
