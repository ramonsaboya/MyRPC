package commons

import (
	"encoding/json"
)

type Marshaller struct{}

func (Marshaller) Marshall(data TempPacket) ([]byte, error) {
	r, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (Marshaller) Unmarshall(data []byte) (TempPacket, error) {
	r := TempPacket{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
