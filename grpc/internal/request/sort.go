package request

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func (r *request) ProcessSort(q query.Query,
	instruction instruction.Instruction) (query.Query, errors.Error) {
	sorts, err := gateway.ProcessSort(r, q, instruction)
	if err != nil {
		return nil, err
	}
	if sorts != nil {
		return sorts, nil
	}
	sorts = instruction.GetDefaultSorts()
	q.WithSub(sorts)
	return q, nil
}
