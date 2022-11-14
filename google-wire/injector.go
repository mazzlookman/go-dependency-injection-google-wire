//go:build wireinject
// +build wireinject

package google_wire

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil
}
