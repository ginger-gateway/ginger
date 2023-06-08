package request

import "github.com/ginger-core/errors"

func (r *request) GetHeader(key string) string {
	return r.Context.GetHeader(key)
}

func (r *request) ProcessHeaders(ref any) errors.Error {
	err := r.ShouldBindHeader(ref)
	if err != nil {
		return errors.Validation(err)
	}

	r.header = ref
	return nil
}

func (r *request) GetHeaders() any {
	return r.header
}
