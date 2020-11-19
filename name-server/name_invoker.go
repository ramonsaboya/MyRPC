package name

import (
	"github.com/ramonsaboya/myrpc/commons"
	"github.com/ramonsaboya/myrpc/server"
)

type NameInvoker struct {
	proxy commons.ClientProxy
}

func NewNamingInvoker(proxy *commons.ClientProxy) *NameInvoker {
	return &NameInvoker{
		proxy: *proxy,
	}
}

func (n *NameInvoker) Invoke() error {
	srh, err := server.NewSRH(n.proxy.Protocol, n.proxy.Host, n.proxy.Port)

	if err != nil {
		return err
	}

	marshaller := commons.Marshaller{}
	nameService := NameServer{
		services: make(map[string]commons.ClientProxy),
	}

	res := commons.TempPacket{}

	var reply interface{}

	for {
		rcvMsgBytes, err := srh.Receive()
		if err != nil {
			return err
		}

		req, err := marshaller.Unmarshall(rcvMsgBytes)
		if err != nil {
			return err
		}
		operation := req.Operation

		switch operation {
		case "Register":
			proxy := req.Proxy
			serviceName := req.ServiceName
			reply, _ = nameService.Register(serviceName, proxy)
		case "Lookup":
			serviceName := req.ServiceName
			reply, _ = nameService.Lookup(serviceName)
		}

		res.Reply = reply
		msgToClientBytes, err := marshaller.Marshall(res)
		if err != nil {
			return err
		}

		err = srh.Send(msgToClientBytes)
		if err != nil {
			return err
		}
	}
}
