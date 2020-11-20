package miop

type Packet struct {
	Hdr Header
	Bd  Body
}

type Header struct {
	MessageType int
}

type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

type RequestHeader struct {
	RequestId string
	ObjectKey int
	Operation string
}

type RequestBody struct {
	Body []interface{}
}

type ReplyHeader struct {
	RequestId string
	Status    int
}

type ReplyBody struct {
	OperationResult interface{}
}
