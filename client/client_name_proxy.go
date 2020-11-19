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

func (n *NameProxy) Lookup(name string) (*commons.ClientProxy, error) {
	req := client.Request{
		Operation:   "Lookup",
		ServiceName: "Calculator",
	}

	requestor, err := client.NewRequestor(&n.proxy)
	if err != nil {
		return nil, err
	}
	res, err := requestor.Invoke(req)

	if err != nil {
		return nil, err
	}

	reply := res.(commons.ClientProxy)

	return &reply, nil
}
