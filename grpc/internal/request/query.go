package request

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func (r *request) ProcessFilters(base query.Query,
	instruction instruction.Instruction) (query.Query, errors.Error) {
	filters, err := gateway.ProcessFilters(r, base, instruction)
	if err != nil {
		return nil, err.WithTrace("gateway.ProcessFilters")
	}
	if filters == nil {
		return base, nil
	}

	return filters, nil
}

func (r *request) GetQuery(key string) (string, bool) {
	v := r.query[key]
	return v, v != ""
}

func (r *request) ProcessQueries(ref any) errors.Error {
	panic("not implemented")
}

func (r *request) GetQueries() any {
	return r.query
}

type queriesGetter interface {
	GetQueries() map[string]string
}

func (r *request) initQueries(req any) errors.Error {
	if getter, ok := req.(queriesGetter); ok {
		r.query = getter.GetQueries()
	}
	return nil
}
