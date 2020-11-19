package nameserver

import (
	"reflect"

	"github.com/ramonsaboya/myrpc/commons"
)

type NameServer struct {
	services map[string]map[string]commons.ClientProxy
}

func (ns *NameServer) Register(name string, protocol string, proxy commons.ClientProxy) (bool, error) {
	if ns.services[protocol] == nil {
		ns.services[protocol] = map[string]commons.ClientProxy{}
	}

	ns.services[protocol][name] = proxy

	return true, nil
}

func (ns *NameServer) Lookup(name string, protocol string) (commons.ClientProxy, error) {
	if ns.services[protocol] == nil {
		return commons.ClientProxy{}, nil
	}

	if reflect.ValueOf(ns.services[protocol][name]).IsNil() {
		return commons.ClientProxy{}, nil
	}

	return ns.services[protocol][name], nil
}

func (ns *NameServer) List(name string, protocol string) (map[string]commons.ClientProxy, error) {
	if ns.services[protocol] == nil {
		return map[string]commons.ClientProxy{}, nil
	}

	return ns.services[protocol], nil
}
