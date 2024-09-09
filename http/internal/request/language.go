package request

import "github.com/ginger-core/gateway"

func (r *request) WithLanguage(language gateway.Language) gateway.Request {
	r.language = language
	return r
}

func (r *request) GetLanguage() gateway.Language {
	return r.language
}
