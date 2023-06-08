package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal"
)

func (r *responder) RespondError(request gateway.Request, err errors.Error) {
	c := request.GetConn().(*gin.Context)
	c.Header(internal.RequestIdHeaderKey, request.GetId())

	if err == nil {
		c.Status(http.StatusExpectationFailed)
		c.Abort()
		request.SetResponded()
		return
	}

	language := request.GetLanguage()
	if language != nil {
		err = gateway.TranslateError(language, err)

		extras := err.GetExtra()
		for _, es := range extras {
			for _, e := range es {
				if ce := e.GetError().(errors.Error); ce != nil {
					e.WithError(gateway.TranslateError(language, ce))
				}
			}
		}
	}
	_ = c.Error(err)
	c.JSON(getStatusCodeByErrType(err.GetType()), err)
	c.Abort()
	request.SetResponded()
}
