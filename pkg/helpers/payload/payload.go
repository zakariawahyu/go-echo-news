package payload

type Payload struct {
	Limit  int
	Offset int
}

func NewPayload() Payload {
	return Payload{
		Limit:  20,
		Offset: 0,
	}
}
