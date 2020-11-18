package nameserver

import "github.com/ramonsaboya/myrpc/commons"

type NameServer struct {
	services map[string]commons.ClientProxy
}

func (ns *NameServer) Register(name string, proxy commons.ClientProxy) (bool, error) {

}

func (ns *NameServer) Lookup(name string) (commons.ClientProxy, error) {

}

func (ns *NameServer) List(name string) (map[string]commons.ClientProxy, error) {

}
