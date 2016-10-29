package mano

// DefaultRoutes 是用于收集默认路由的集合
var DefaultRouters = &RouterGroup{}

// Get registers a new GET request handle and middleware with the given pattern.
func Get(pattern string, handler Handler, middlewares ...string) *Router {
	return DefaultRouters.Get(pattern, handler, middlewares...)
}
