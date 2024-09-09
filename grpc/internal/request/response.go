package request

func (r *request) SetResponded() {
	r.hasResponded = true
}

func (r *request) HasResponded() bool {
	return r.hasResponded
}
