package mano

import (
	"container/list"
	"net/url"
	"regexp"
)

type RouteData struct {
	route      *Route
	matchedUrl string
	handler    HandlerFunc
	values     map[string]string
}

type Route struct {
	pattern *regexp.Regexp
	handler HandlerFunc
}

type RouteTable struct {
	routes *list.List
}

func (rt *RouteTable) Register(pattern *regexp.Regexp, handler HandlerFunc, segments int, endsWildcard bool) {
	route := &Route{
		pattern: pattern,
		handler: handler,
	}
	//TODO:路由的优先级
	rt.routes.PushBack(route)
}

//http://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-golang-regular-expressions

// Match 匹配
func (rt *RouteTable) Match(url *url.URL) (*RouteData, bool) {
	//url := ""
	var route *Route
	elem := rt.routes.Front()
	for elem != nil {
		route, _ = elem.Value.(*Route)
		elem = elem.Next()
		if route == nil {
			continue
		}
		if route.pattern.MatchString(url.Path) {
			return &RouteData{
				route:      route,
				matchedUrl: url.Path,
				handler:    route.handler, //TODO
			}, true
		}
	}
	return nil, false
}
