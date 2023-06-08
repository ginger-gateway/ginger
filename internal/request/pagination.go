package request

import (
	"strconv"

	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/query"
)

func (r *request) GetPagination() (int, int) {
	var size, page int
	sizeQ, ok := r.Context.GetQuery("size")
	if ok {
		size, _ = strconv.Atoi(sizeQ)
	}
	pageQ, ok := r.Context.GetQuery("page")
	if ok {
		page, _ = strconv.Atoi(pageQ)
	}
	if size <= 0 {
		size = 30
	}
	if page <= 0 {
		page = 1
	}
	return size, page
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
