// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package google_wire

// Injectors from injector.go:

func InitializedService() (*SimpleService, error) {
	simpleRepository := NewSimpleRepository()
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}
