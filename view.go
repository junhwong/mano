package mano

import (
	"encoding/json"
)

// View for a web interaction.
// Implementations are responsible for rendering content, and exposing the model.
type View interface {
	// ContentType returns HTTP Content-type
	ContentType() string
	// Render the view given the specified model.
	Render(ctx Context) error
}

type ActionView struct {
	local       string
	contentType string
}

func (v *ActionView) ContentType() string {
	if v.contentType == "" {
		return "text/plain; charset=UTF-8"
	}
	return v.contentType
}

type ContentView struct {
	*ActionView
	Content string
}

// func (ar *ContentView) ContentType() string {
// 	if ar.contentType == "" {
// 		return "text/plain; charset=UTF-8"
// 	}
// 	return ar.contentType
// }

func (ar *ContentView) Render(ctx Context) error {
	w := ctx.Response().Writer()
	// w.Header().Set("Content-Type", ar.ContentType())
	_, err := w.Write([]byte(ar.Content))
	return err
}

type JsonView struct {
	*ActionView
	Data interface{}
}

// func (v *JsonView) ContentType() string {
// 	if v.contentType == "" {
// 		return "application/json; charset=UTF-8"
// 	}
// 	return v.contentType
// }

func (v *JsonView) Render(ctx Context) error {
	content, err := json.Marshal(v.Data)
	if err != nil {
		return err
	}
	w := ctx.Response().Writer()
	// w.Header().Set("Content-Type", v.ContentType())
	_, err = w.Write([]byte(content))
	return err
}
