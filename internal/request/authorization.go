package request

import "github.com/ginger-core/gateway"

func (r *request) IsAuthenticated() bool {
	return r.authorization != nil && r.authorization.GetApplicantId() != nil
}

func (r *request) SetAuthorization(authorization gateway.Authorization) {
	r.authorization = authorization
}

func (r *request) WithAuthorization(auth gateway.Authorization) gateway.Request {
	r.authorization = auth
	return r
}

func (r *request) GetAuthorization() gateway.Authorization {
	return r.authorization
}
