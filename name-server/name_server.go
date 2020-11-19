package name

import (
	"reflect"

	"github.com/ramonsaboya/myrpc/commons"
)

type NameServer struct {
	services map[string]commons.ClientProxy
}

func (ns *NameServer) Register(name string, proxy commons.ClientProxy) (bool, error) {
	ns.services[name] = proxy

	return true, nil
}

func (ns *NameServer) Lookup(name string) (commons.ClientProxy, error) {
	if reflect.ValueOf(ns.services[name]).IsNil() {
		return commons.ClientProxy{}, nil
	}

	return ns.services[name], nil
}

func (ns *NameServer) List(name string) (map[string]commons.ClientProxy, error) {
	return ns.services, nil
}
