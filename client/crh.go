package client

import (
	"net"
	"unsafe"

	"github.com/ramonsaboya/myrpc/commons"
)

// in a real world scenario, this connection cache would require
// some kind of eviction policy, but this is okay for our use case
var connCache = make(map[string]*net.Conn)

type CRH struct {
	address string

	protocol   commons.Protocol
	bufferSize int
}

func newCRH(protocol commons.Protocol, address string, bufferSize int) (*CRH, error) {
	return &CRH{
		address:    address,
		protocol:   protocol,
		bufferSize: bufferSize,
	}, nil
}

func createConnection(protocol commons.Protocol, address string) (*net.Conn, error) {
	conn, err := net.Dial(string(protocol), address)
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

func (crh *CRH) getConnection() (*net.Conn, error) {
	conn, ok := connCache[crh.address]
	if !ok {
		conn, err := createConnection(crh.protocol, crh.address)
		if err != nil {
			return nil, err
		}
		connCache[crh.address] = conn
	}

	return conn, nil
}

func (crh *CRH) SendReceive(data []byte) ([]byte, error) {
	conn, err := crh.getConnection()
	if err != nil {
		return nil, err
	}

	(*conn).Write(data)

	buf := make([]byte, unsafe.Sizeof(crh.bufferSize))
	_, err = (*conn).Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
