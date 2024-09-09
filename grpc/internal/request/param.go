package request

func (r *request) GetParam(key string) string {
	return r.params[key]
}

type paramsGetter interface {
	GetParams() map[string]string
}

func (r *request) initParams(req any) {
	if getter, ok := req.(paramsGetter); ok {
		r.params = getter.GetParams()
	}
}
