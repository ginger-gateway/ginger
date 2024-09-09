package request

func (r *request) GetParam(key string) string {
	return r.Context.Param(key)
}
