package google_wire

type FooBarService struct {
	*FooService
	*BarService
}

func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}

type FoobarStruct struct {
	*FooStruct
	*BarStruct
}
