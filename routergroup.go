package mano

type RouterGroup struct {
	Router
	items []*Router
}

func (group *RouterGroup) add(method HttpMethod, pattern string, handler Handler, middlewares []string) *Router {
	router := &Router{
		method:      method,
		pattern:     pattern,
		handler:     handler,
		middlewares: middlewares,
	}
	group.items = append(group.items, router)
	return router
}

func (group *RouterGroup) Get(pattern string, handler Handler, middlewares ...string) *Router {
	return group.add(GET, pattern, handler, middlewares)
}
