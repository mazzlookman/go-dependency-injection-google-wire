//go:build wireinject
// +build wireinject

package google_wire

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService(foo Foo, bar Bar) *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var sayHelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedSayHelloService() *SayHelloService {
	wire.Build(sayHelloSet, NewSayHelloService)
	return nil
}

func InitializedFooBarServiceStruct(foo Foo, bar Bar) *FooBarService {
	wire.Build(fooSet, wire.Struct(new(FooBarService), "FooService"))
	return nil
}

var fooValue = &FooStruct{}
var barValue = &BarStruct{}

func InitializedFooBarStructUsingValue() *FoobarStruct {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FoobarStruct), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
