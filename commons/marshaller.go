package commons

import (
	"encoding/json"

	"github.com/ramonsaboya/myrpc/miop"
)

type Marshaller struct{}

func (Marshaller) Marshall(data miop.Packet) ([]byte, error) {
	r, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (Marshaller) Unmarshall(data []byte) (miop.Packet, error) {
	r := miop.Packet{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
