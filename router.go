package djgo

import "net/http"

type controllerInfo struct{}

type userHandler struct{}

type ControllerRegistor struct {
	routers      []*controllerInfo
	userHandlers map[string]*userHandler
}

// AutoRoute
func (p *ControllerRegistor) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
}

func NewControllerRegistor() *ControllerRegistor {
	return &ControllerRegistor{routers: make([]*controllerInfo, 0), userHandlers: make(map[string]*userHandler)}
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {
}
