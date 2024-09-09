package grpc

type config struct {
	Logger struct {
		SkipPaths []string
	}
	ListenAddr string
}

func (c *config) initialize() {
}
