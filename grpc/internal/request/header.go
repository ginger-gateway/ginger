package request

import "github.com/ginger-core/errors"

func (r *request) GetHeader(key string) string {
	return r.headers[key]
}

func (r *request) ProcessHeaders(ref any) errors.Error {
	panic("not implemented")
}

func (r *request) GetHeaders() any {
	return r.headers
}

type headersGetter interface {
	GetHeaders() map[string]string
}

func (r *request) initHeaders(req any) errors.Error {
	if getter, ok := req.(headersGetter); ok {
		r.headers = getter.GetHeaders()
	}
	return nil
}
