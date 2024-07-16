package router

import "sync"

type Route struct {
	Address       string
	IsMaintenance bool
}

type Routes struct {
	sync.Mutex
	ActiveAddress []string
	KnownRoutes   map[string]Route
}

func (r *Routes) GetRoute(n int64) *Route {
	len := int64(len(r.ActiveAddress))
	if len == 0 {
		return nil
	}

	r.Lock()
	address := r.ActiveAddress[n%len]
	r.Unlock()

	if route, ok := r.KnownRoutes[address]; ok {
		return &route
	}
	return nil
}

func (r *Routes) AddRoute(address string) {
	r.Lock()
	defer r.Unlock()

	r.KnownRoutes[address] = Route{Address: address, IsMaintenance: false}
	r.ActiveAddress = append(r.ActiveAddress, address)
}
