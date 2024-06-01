package actor

type Request struct {
	Data []byte
}

func newSignRequest(data []byte) *Request {
	return &Request{
		Data: data,
	}
}

type Signature struct {
	Data []byte
}

func newSignResponse(data []byte) *Signature {
	return &Signature{
		Data: data,
	}
}
