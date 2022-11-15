package test

import (
	"github.com/stretchr/testify/assert"
	google_wire "go-salaries-app/google-wire"
	"testing"
)

func TestSimpleServiceFailed(t *testing.T) {
	simpleService, err := google_wire.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := google_wire.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}

func TestDatabaseRepository(t *testing.T) {
	databaseRepository := google_wire.InitializedDatabaseRepository()
	assert.Equal(t, "PostgreSQL", databaseRepository.DatabasePostgreSQL.Name)
	assert.Equal(t, "MongoDB", databaseRepository.DatabaseMongoDB.Name)
}

func TestFooBarService(t *testing.T) {
	foo := google_wire.Foo("I'am a Foo")
	bar := google_wire.Bar("I'am a Bar")
	fooBarService := google_wire.InitializedFooBarService(foo, bar)
	assert.Equal(t, "I'am a Foo", string(fooBarService.FooService.Name))
	assert.Equal(t, "I'am a Bar", string(fooBarService.BarService.Name))
}

func TestSayHelloService(t *testing.T) {
	helloService := google_wire.InitializedSayHelloService()
	hello := helloService.SayHello.Hello("iamaqibmoh")
	assert.Equal(t, "Hello iamaqibmoh", hello)
}

func TestFooBarServiceStruct(t *testing.T) {
	foo := google_wire.Foo("I'am a Foo")
	bar := google_wire.Bar("I'am a Bar")
	fooBarServiceStruct := google_wire.InitializedFooBarServiceStruct(foo, bar)
	assert.Equal(t, "I'am a Foo", string(fooBarServiceStruct.FooService.Name))
}

func TestFooBarStructUsingValue(t *testing.T) {
	foobarStruct := google_wire.InitializedFooBarStructUsingValue()
	foobarStruct.FooStruct.Name = "Foo"
	foobarStruct.BarStruct.Name = "Bar"
	assert.Equal(t, "Foo", foobarStruct.FooStruct.Name)
	assert.Equal(t, "Bar", foobarStruct.BarStruct.Name)
}

func TestConfiguration(t *testing.T) {
	conf := google_wire.InitializedConfiguration()
	assert.Equal(t, "iamaqibmoh", conf.Name)
}

func TestConnection(t *testing.T) {
	connection, cleanup := google_wire.InitializedConnection("Database")
	assert.NotNil(t, connection)
	cleanup()
}
