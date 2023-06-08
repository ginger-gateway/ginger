package request

import "github.com/ginger-core/errors"

func (r *request) ProcessBody(ref any) errors.Error {
	err := r.ShouldBind(ref)
	if err != nil {
		return errors.Validation(err)
	}

	r.body = ref
	return nil
}

func (r *request) GetBody() any {
	return r.body
}
