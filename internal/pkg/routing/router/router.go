package router

import (
	"sync/atomic"
)

var (
	ops atomic.Uint32
	routes Routes
)

func Init() {
	routes = Routes{
		ActiveAddress: make([]string, 0),
		KnownRoutes: make(map[string]Route),
	}
}

func GetAddress() string {
	try := 0
	for try < 5 {
		current := int64(ops.Add(1))
		route := routes.GetRoute(current)
		if route != nil && !route.IsMaintenance {
			return route.Address
		}
		
		try = try + 1
	}
	return ""
}

func Register(address string) {
	routes.AddRoute(address)
}