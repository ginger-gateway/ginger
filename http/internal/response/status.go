package response

import (
	"net/http"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
)

var statusMap = map[gateway.Status]int{
	gateway.StatusUnknown:           http.StatusNotImplemented,
	gateway.StatusOK:                http.StatusOK,
	gateway.StatusCreated:           http.StatusCreated,
	gateway.StatusNoContent:         http.StatusNoContent,
	gateway.StatusMovedPermanently:  http.StatusMovedPermanently,
	gateway.StatusFound:             http.StatusNotFound,
	gateway.StatusPermanentRedirect: http.StatusPermanentRedirect,
	gateway.StatusTemporaryRedirect: http.StatusTemporaryRedirect,
	gateway.StatusBadRequest:        http.StatusBadRequest,
	gateway.StatusDuplicate:         http.StatusConflict,
}

func getStatusCode(status gateway.Status) int {
	if code := statusMap[status]; code > 0 {
		return code
	}
	return http.StatusNotImplemented
}

var errTypeStatusMap = map[errors.Type]int{
	errors.TypeInternal:         500,
	errors.TypeValidation:       400,
	errors.TypeNotFound:         404,
	errors.TypeUnauthorized:     401,
	errors.TypeForbidden:        403,
	errors.TypeTooManyRequests:  429,
	errors.TypeFailedDependency: 424,
	errors.TypeTooEarly:         425,
	errors.TypeExpired:          408,
	errors.TypeDuplicate:        409,
}

func getStatusCodeByErrType(t errors.Type) int {
	if code := errTypeStatusMap[t]; code > 0 {
		return code
	}
	return http.StatusInternalServerError
}
