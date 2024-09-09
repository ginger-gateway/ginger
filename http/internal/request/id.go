package request

import "github.com/ginger-core/gateway"

func (r *request) WithId(id string) gateway.Request {
	r.id = id
	return r
}

func (r *request) GetId() string {
	return r.id
}
