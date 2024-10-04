package request

import (
	"context"

	"github.com/ginger-core/gateway"
	"github.com/google/uuid"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type request struct {
	// Context equivalent to connection
	Context any
	// context as connection context to handle timeout, cancel, etc.
	context context.Context
	// authorization of applicant
	authorization gateway.Authorization
	// id request unique id
	id string
	// query is processed query of request
	query map[string]string
	// params is processed params of request
	params map[string]string
	// headers is processed headers of request
	headers map[string]string
	// body is processed body of request
	body any
	// hasResponded determines if it has already responded to client or not
	hasResponded bool
	// language current language of request
	language gateway.Language
}

func New(ctx any, langBundle *i18n.Bundle, req any) gateway.Request {
	r := &request{
		Context: ctx,
		body:    req,
	}
	if r.id == "" {
		uid, _ := uuid.NewRandom()
		r.id = uid.String()
	}
	r.initHeaders(req)
	r.initQueries(req)
	r.initParams(req)
	if langBundle != nil {
		acceptLanguage := r.GetHeader("Accept-Language")
		r.language = gateway.NewLanguage(langBundle, acceptLanguage, "EN")
	}
	return r
}
