package client

import (
	"github.com/mitchellh/mapstructure"
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
	params := make([]interface{}, 1)
	params[0] = name
	req := Request{
		Operation: "Lookup",
		Params:    params,
	}

	requestor, err := NewRequestor(&n.proxy)
	if err != nil {
		return nil, err
	}
	res, err := requestor.Invoke(req)

	if err != nil {
		return nil, err
	}
	resMap := res.(map[string]interface{})
	proxy := commons.ClientProxy{}

	mapstructure.Decode(resMap, &proxy)
	return &proxy, nil
}
