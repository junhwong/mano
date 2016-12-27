package mano

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/junhwong/mano/logs"
)

type RouteHandler struct {
	app    *Application
	prefix []string
}

func (handler *RouteHandler) Init(app *Application) error {
	handler.app = app
	return nil
}

func (handler *RouteHandler) Handle(writer http.ResponseWriter, request *http.Request) (complated bool, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = errors.New(fmt.Sprint(r))
			}
		}
	}()

	routeData, matched := handler.app.routeTable.Match(ParseHttpMethod(request.Method), request.URL)
	if !matched {
		complated = false
		return
	}

	ctx := newRequestContext(handler.app, request, writer, routeData)

	ch := &middlewareChan{
		app:         handler.app,
		handler:     routeData.entry.handler,
		index:       0,
		middlewares: routeData.entry.middlewares,
	}
	ctx.Data("lang", handler.app.lang) //设置默认语言资源到上下文
	result := ch.exec(ctx)
	view, ok := result.(View)
	if ok {

	} else if s, ok := result.(string); ok {
		if strings.HasPrefix(s, "view:") {
			err = handler.app.viewEngine.Render(ctx.data, s[5:], writer)
			if err != nil {
				logs.Debug(err)
			}
			return
			//view=ctx.View(s[5:])
		} else {
			view = ctx.Content(s)
		}
	} else {
		panic("todo")
	}

	contentType := view.ContentType()
	if contentType == "" {
		contentType = "text/plain; charset=UTF-8"
	}
	writer.Header().Set("Content-Type", contentType)
	view.Render(ctx)
	return
}

func URLRouting(prefix ...string) *RouteHandler {
	return &RouteHandler{
		prefix: prefix,
	}
}
