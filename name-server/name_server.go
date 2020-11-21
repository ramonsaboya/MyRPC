package name

import (
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
	proxy := ns.services[name]
	return proxy, nil
}

func (ns *NameServer) List(name string) (map[string]commons.ClientProxy, error) {
	return ns.services, nil
}
