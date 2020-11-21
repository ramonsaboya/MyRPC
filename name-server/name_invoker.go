package name

import (
	"github.com/mitchellh/mapstructure"
	"github.com/ramonsaboya/myrpc/commons"
	"github.com/ramonsaboya/myrpc/miop"
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

	res := miop.Packet{}

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
		operation := req.Bd.ReqHeader.Operation

		switch operation {
		case "Register":
			proxy := commons.ClientProxy{}
			mapstructure.Decode(req.Bd.ReqBody.Body[1], &proxy)
			serviceName := req.Bd.ReqBody.Body[0].(string)
			reply, _ = nameService.Register(serviceName, proxy)
		case "Lookup":
			serviceName := req.Bd.ReqBody.Body[0].(string)
			reply, _ = nameService.Lookup(serviceName)
		}

		repHeader := miop.ReplyHeader{RequestId: req.Bd.ReqHeader.RequestId, Status: 200}
		repBody := miop.ReplyBody{OperationResult: reply}
		header := miop.Header{MessageType: commons.MIOPREQUEST}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		res = miop.Packet{Hdr: header, Bd: body}

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
