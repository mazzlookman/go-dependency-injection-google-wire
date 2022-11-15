package google_wire

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type Foo string

type FooService struct {
	*FooRepository
	Name Foo
}

func NewFooService(fooRepository *FooRepository, name Foo) *FooService {
	return &FooService{
		FooRepository: fooRepository,
		Name:          name,
	}
}
