package request

import "github.com/ginger-core/errors"

func (r *request) ProcessBody(ref any) errors.Error {
	r.body = ref
	return nil
}

func (r *request) GetBody() any {
	return r.body
}
