package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal"
)

func (r *responder) Respond(request gateway.Request,
	status gateway.Status, response any) {
	request.SetResponded()
	c := request.GetConn().(*gin.Context)
	c.Header(internal.RequestIdHeaderKey, request.GetId())

	switch status {
	case gateway.StatusMovedPermanently:
		c.Redirect(http.StatusMovedPermanently, response.(string))
		return
	case gateway.StatusFound:
		c.Redirect(http.StatusFound, response.(string))
		return
	case gateway.StatusPermanentRedirect:
		c.Redirect(http.StatusPermanentRedirect, response.(string))
		return
	case gateway.StatusTemporaryRedirect:
		c.Redirect(http.StatusTemporaryRedirect, response.(string))
		return
	case gateway.StatusUnknown:
		switch c.Request.Method {
		case http.MethodPost:
			if response == nil {
				status = gateway.StatusNoContent
				break
			}
			status = gateway.StatusCreated
		case http.MethodGet:
			status = gateway.StatusOK
		case http.MethodPut:
			if response == nil {
				status = gateway.StatusNoContent
				break
			}
			status = gateway.StatusOK
		case http.MethodDelete:
			if response == nil {
				status = gateway.StatusNoContent
				break
			}
			status = gateway.StatusOK
		}
	}
	switch v := response.(type) {
	case gateway.Response:
		c.Header("Content-Type", parseContentType(v.GetContentType()))
		for k, v := range v.GetHeaders() {
			c.Header(k, v)
		}
		c.Data(getStatusCode(status),
			parseContentType(v.GetContentType()),
			v.GetBody().Bytes())
	default:
		if response == nil {
			c.Status(getStatusCode(status))
		} else {
			c.JSON(getStatusCode(status), response)
		}
	}
}
