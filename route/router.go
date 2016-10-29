package route

// import "github.com/junhwong/mano"

// //Context is an alias for mano.Context.
// type Context mano.Context

// type HandlerFunc func(Context) interface{}

// //Router is defined URL mapping
// type router struct {
// 	path        string
// 	middlewares []string
// 	handler     HandlerFunc
// }

// func Group() *router {
// 	return nil
// }

// func newRouter(method string, path string, handler HandlerFunc, middlewares []string) mano.Route {
// 	//保存路由定义的位置，方便调试
// 	// pc, file, line, ok := runtime.Caller(1)
// 	// log.Println(pc)
// 	// log.Println(file)
// 	// log.Println(line)
// 	// log.Println(ok)
// 	// f := runtime.FuncForPC(pc)
// 	// log.Println(f.Name())

// 	router := &router{
// 		path:        path,
// 		handler:     handler,
// 		middlewares: middlewares,
// 	}

// 	return mano.AppendRoute(router)
// }

// //*************** mano.Route 接口 ***************

// func (r *router) Handle(ctx mano.Context) interface{} {
// 	return r.handler(Context(ctx))
// }

// func (r *router) Ensure(mano.Application) (string, mano.Route, error) {
// 	return "", r, nil
// }

// //*************** 默认API ***************

// //Get registers a new GET request handle and middleware with the given path.
// func Get(path string, handler HandlerFunc, middlewares ...string) mano.Route {
// 	return newRouter("GET", path, handler, middlewares)
// }

// //Post registers a new POST request handle and middleware with the given path.
// func Post(path string, handler HandlerFunc, middlewares ...string) mano.Route {
// 	return newRouter("POST", path, handler, middlewares)
// }

// //Put registers a new PUT request handle and middleware with the given path.
// func Put(path string, handler HandlerFunc, middlewares ...string) mano.Route {
// 	return newRouter("PUT", path, handler, middlewares)
// }

// //Delete registers a new DELETE request handle and middleware with the given path.
// func Delete(path string, handler HandlerFunc, middlewares ...string) mano.Route {
// 	return newRouter("DELETE", path, handler, middlewares)
// }

// //Options registers a new OPTIONS request handle and middleware with the given path.
// func Options(path string, handler HandlerFunc, middlewares ...string) mano.Route {
// 	return newRouter("OPTIONS", path, handler, middlewares)
// }
