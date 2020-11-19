package client

import (
	"github.com/ramonsaboya/myrpc/commons"
)

type Requestor struct {
	proxy     *commons.ClientProxy
	RequestId int
}

type Request struct {
	Operation string
	Params    []interface{}
}

func NewRequestor(proxy *commons.ClientProxy) (*Requestor, error) {
	return &Requestor{
		proxy:     proxy,
		RequestId: 0,
	}, nil
}

func (r *Requestor) Invoke(inv Request) (interface{}, error) {
	marshaller := commons.Marshaller{}
	crh, err := NewCRH(r.proxy.Protocol, r.proxy.Host, r.proxy.Port)
	if err != nil {
		return nil, err
	}

	reqHeader := commons.RequestHeader{RequestId: r.RequestId, ObjectKey: 1, Operation: inv.Operation}
	reqBody := commons.RequestBody{Body: inv.Params}
	header := commons.Header{MessageType: commons.TEMPREQUEST}
	body := commons.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	tempPacketRequest := commons.TempPacket{Hdr: header, Bd: body}

	msgToClientBytes, err := marshaller.Marshall(tempPacketRequest)
	if err != nil {
		return nil, err
	}
	msgFromServerBytes, err := crh.SendReceive(msgToClientBytes)
	if err != nil {
		return nil, err
	}
	tempPacketReply, err := marshaller.Unmarshall(msgFromServerBytes)
	if err != nil {
		return nil, err
	}

	return tempPacketReply.Bd.RepBody.OperationResult, nil
}
