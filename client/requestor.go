package client

import (
	"github.com/ramonsaboya/myrpc/commons"
)

type Requestor struct {
	proxy *commons.ClientProxy
}

type Request struct {
	Operation string

	Params      []interface{}
	Proxy       commons.ClientProxy
	ServiceName string
}

func NewRequestor(proxy *commons.ClientProxy) (*Requestor, error) {
	return &Requestor{
		proxy: proxy,
	}, nil
}

func (r *Requestor) Invoke(inv Request) (interface{}, error) {
	marshaller := commons.Marshaller{}
	crh, err := NewCRH(r.proxy.Protocol, r.proxy.Host, r.proxy.Port)
	if err != nil {
		return nil, err
	}

	msgToClientBytes, err := marshaller.Marshall(commons.TempPacket{
		Operation:   inv.Operation,
		Params:      inv.Params,
		Reply:       make([]interface{}, 0),
		Proxy:       inv.Proxy,
		ServiceName: inv.ServiceName,
	})
	if err != nil {
		return nil, err
	}

	msgFromServerBytes, err := crh.SendReceive(msgToClientBytes)
	if err != nil {
		return nil, err
	}
	rawReply, err := marshaller.Unmarshall(msgFromServerBytes)
	if err != nil {
		return nil, err
	}

	return rawReply.Reply, nil
}
