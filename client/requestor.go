package client

import (
	"math/rand"
	"time"

	"github.com/ramonsaboya/myrpc/commons"
	"github.com/ramonsaboya/myrpc/miop"
)

type Requestor struct {
	proxy *commons.ClientProxy
}

type Request struct {
	Operation string
	Params    []interface{}
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func NewRequestor(proxy *commons.ClientProxy) (*Requestor, error) {
	rand.Seed(time.Now().UnixNano())
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

	requestId := randomString(32)
	reqHeader := miop.RequestHeader{RequestId: requestId, ObjectKey: 1, Operation: inv.Operation}
	reqBody := miop.RequestBody{Body: inv.Params}
	header := miop.Header{MessageType: commons.MIOPREQUEST}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	packetRequest := miop.Packet{Hdr: header, Bd: body}

	msgToClientBytes, err := marshaller.Marshall(packetRequest)
	if err != nil {
		return nil, err
	}
	msgFromServerBytes, err := crh.SendReceive(msgToClientBytes)
	if err != nil {
		return nil, err
	}
	packetReply, err := marshaller.Unmarshall(msgFromServerBytes)
	if err != nil {
		return nil, err
	}

	return packetReply.Bd.RepBody.OperationResult, nil
}
