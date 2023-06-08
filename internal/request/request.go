package request

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger/internal"
	"github.com/google/uuid"
)

type request struct {
	// Context gin context equivalent to connection
	*gin.Context
	// context as connection context to handle timeout, cancel, etc.
	context context.Context
	// authorization of applicant
	authorization gateway.Authorization
	// id request unique id
	id string
	// query is processed query of request
	query any
	// body is processed body of request
	body any
	// header is processed header of request
	header any
	// hasResponded determines if it has already responded to client or not
	hasResponded bool
	// language current language of request
	language gateway.Language
}

func New(c *gin.Context, language gateway.Language) gateway.Request {
	r := &request{
		Context:  c,
		language: language,
		id:       c.GetHeader(internal.RequestIdHeaderKey),
	}
	if r.id == "" {
		uid, _ := uuid.NewRandom()
		r.id = uid.String()
	}
	return r
}
