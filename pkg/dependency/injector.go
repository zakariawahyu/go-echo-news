//go:build wireinject
// +build wireinject

package dependency

import "github.com/google/wire"

// simple.go
func InitialSimpleServices(isErr bool) (*SimpleServices, error) {
	wire.Build(NewSimpleRepository, NewSimpleServices)
	return nil, nil
}

// multi_binding.go
func InitialDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)
	return nil
}

// provider_set.go
// foo.go
// bar.go
var barSet = wire.NewSet(NewBarRepository, NewBarServices)
var fooSet = wire.NewSet(NewFooRepository, NewFooServices)

func InitialFooBarServices() *FooBarServices {
	wire.Build(barSet, fooSet, NewFooBarServices)
	return nil
}

// binding_interface.go
// contoh ijector yang salah
//func InitialHelloServices() *SimpleServices {
//	wire.Build(NewSayHelloImpl, NewHelloServices)
//	return nil
//}

// jadi panggil function impl dulu, terus bind interface dan struct kontraknya
var helloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitialHelloServices() *HelloServices {
	wire.Build(helloSet, NewHelloServices)
	return nil
}

// struct_provider.go
var FooBarSet = wire.NewSet(NewFoo, NewBar)

func InitialFooBar() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "Foo", "Bar")) // atau * untuk semua field
	return nil
}

// Binding Value
var FooBarValueSet = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitialFooBarBindingValue() *FooBar {
	wire.Build(FooBarValueSet, wire.Struct(new(FooBar), "*"))
	return nil
}
