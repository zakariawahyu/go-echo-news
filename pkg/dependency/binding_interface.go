package dependency

type SayHello interface {
	Hello(name string) string
}

type HelloServices struct {
	SayHello SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloServices(hello SayHello) *HelloServices {
	return &HelloServices{
		SayHello: hello,
	}
}
