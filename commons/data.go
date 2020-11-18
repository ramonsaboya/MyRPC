package commons

type TempPacket struct {
	Operation string
	Params    []interface{}
	Reply     interface{}
}
