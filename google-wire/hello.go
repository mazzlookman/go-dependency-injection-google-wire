package google_wire

type SayHello interface {
	Hello(name string) string
}

type SayHelloImpl struct {
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

type SayHelloService struct {
	SayHello
}

func NewSayHelloService(sayHello SayHello) *SayHelloService {
	return &SayHelloService{SayHello: sayHello}
}
