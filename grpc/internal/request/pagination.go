package request

import (
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func (r *request) GetPagination() (int, int) {
	return 30, 1
}

func (r *request) ProcessPagination(q query.Query,
	instruction instruction.Instruction) query.Query {
	size, page := r.GetPagination()
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}
	q = query.NewPagination(q).
		WithSize(size).
		WithPage(page)
	return q
}
