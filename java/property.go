package java

import "sync"

type FakeProperty interface {
	Get(name string) (value any, ok bool)
	Set(name string, value any)
}

type fakeProperty struct {
	data sync.Map
}

func (prop *fakeProperty) Get(name string) (any, bool) {
	return prop.data.Load(HashCode(name))
}

func (prop *fakeProperty) Set(name string, value any) {
	prop.data.Store(HashCode(name), value)
}
