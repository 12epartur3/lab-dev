package main

import(
	"fmt"
)


type ConfigRegistry interface {
	BindProto(key string, proto interface{}) error
	Get(key string) (interface{}, error)
}

type SpexRegistry struct {
	registry ConfigRegistry
}

type SpexConfigRegistry struct {
	name string
	version int
}

func(scr *SpexConfigRegistry) BindProto(key string, proto interface{}) error {
	return nil
}

func(scr *SpexConfigRegistry) Get(key string) (interface{}, error) {
	var i interface{}
	return i, nil
}

var _ ConfigRegistry = (*SpexConfigRegistry)(nil)

func main() {
	spexRegistry := SpexRegistry{
		registry: &SpexConfigRegistry{},
	}
	spexRegistry.registry.Get("123")
	fmt.Printf("hello\n")
}
