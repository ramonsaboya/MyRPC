package server

import (
	"github.com/ramonsaboya/myrpc/commons"
)

type CalculatorInvoker struct {
	proxy commons.ClientProxy
}

func NewCalculatorInvoker(proxy *commons.ClientProxy) *CalculatorInvoker {
	return &CalculatorInvoker{
		proxy: *proxy,
	}
}

func (c *CalculatorInvoker) Invoke() error {
	srh, err := NewSRH(c.proxy.Protocol, c.proxy.Host, c.proxy.Port)
	if err != nil {
		return err
	}
	marshaller := commons.Marshaller{}
	calculator := Calculator{}
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
		case "EquationRoots":
			_a := int(req.Params[0].(float64))
			_b := int(req.Params[1].(float64))
			_c := int(req.Params[2].(float64))
			reply = calculator.EquationRoots(_a, _b, _c)
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
