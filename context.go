package mano

import (
	"net/http"

	"github.com/junhwong/mano/utils"
)

// Context Encapsulates all HTTP-specific information about an individual HTTP request.
type Context interface {
	App() *Application

	//Response 返回响应写入器
	Response() *Response

	Request() *Request

	// Attr Gets or sets the HttpApplication object for the current HTTP request.
	Attr(name string, value ...interface{}) interface{}

	// Data Gets or sets the Context object for the current HTTP request.
	Data(name string, value ...interface{}) interface{}

	//PathValue returns a route path data with given name
	PathValue(name string) (value string, ok bool)

	Param(name string) string

	Content(content string, contentType ...string) View
	JSON(data interface{}, contentType ...string) View
}

type context struct {
	response  *Response
	request   *Request
	app       *Application
	routeData *RouteData
	attrs     utils.Attribute
	data      utils.Attribute
	local     string
}

func newRequestContext(app *Application, request *http.Request, writer http.ResponseWriter, routeData *RouteData) *context {
	return &context{
		response: &Response{
			Response: request.Response,
			writer:   writer,
		},
		request: &Request{
			Request: request,
		},
		app:       app,
		routeData: routeData,
		attrs:     make(utils.Attribute),
		data:      make(utils.Attribute),
	}
}

func (ctx *context) Content(content string, contentType ...string) View {

	view := &ContentView{
		ActionView: &ActionView{
			local: ctx.local,
		},
		Content: content,
	}
	if len(contentType) > 0 {
		view.contentType = contentType[0]
	}
	return view
}

func (ctx *context) JSON(data interface{}, contentType ...string) View {
	view := &JsonView{
		ActionView: &ActionView{
			local: ctx.local,
		},
		Data: data,
	}
	if len(contentType) > 0 {
		view.contentType = contentType[0]
	} else {
		view.contentType = "application/json; charset=UTF-8"
	}
	return view
}

func (ctx *context) App() *Application {
	return ctx.app
}

func (ctx *context) Response() *Response {
	return ctx.response
}
func (ctx *context) Request() *Request {
	return ctx.request
}

func (ctx *context) Attr(name string, value ...interface{}) interface{} {
	return ctx.attrs.Item(name, value...)
}

func (ctx *context) Data(name string, value ...interface{}) interface{} {
	return ctx.data.Item(name, value...)
}

func (ctx *context) PathValue(name string) (value string, ok bool) {
	value = ""
	ok = false
	return
}

func (ctx *context) Param(name string) string {
	return "nil"
}

// func (ctx *context) Log(format string, value ...interface{}) {
// 	return ctx.engine.Logger().Debug("")
// }
