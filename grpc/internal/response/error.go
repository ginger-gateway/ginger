package response

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

func (r *responder) RespondError(request gateway.Request, err errors.Error) {
	panic("not implemented")
}
