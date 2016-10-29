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
			//logs.Error(r)
			err = errors.New(fmt.Sprint(r))
			// log.Debug("%s\n", r)
			// ctx.writer.WriteHeader(http.StatusInternalServerError)
		}
	}()

	routeData, matched := handler.app.routeTable.Match(request.URL)
	if !matched {
		complated = false
		return
	}

	ctx := newRequestContext(handler.app, request, writer, routeData)

	ch := &middlewareChan{
		app:         handler.app,
		handler:     routeData.handler,
		index:       0,
		middlewares: []Middleware{
		//mid, test,
		},
	}

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
