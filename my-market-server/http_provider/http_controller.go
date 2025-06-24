package httpprovider

import (
	"context"
	"net/http"
)

type GenericControllerAdapter[T any] struct {
	Controller Controller[T]
}

type Controller[T any] struct {
	Path    string
	Method  string
	Handler func(r HttpResponse[T])
}

type IController interface {
	GetPath() string
	GetMethod() string
	Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}

func (c GenericControllerAdapter[T]) GetPath() string {
	return c.Controller.Path
}

func (c GenericControllerAdapter[T]) GetMethod() string {
	return c.Controller.Method
}

func (c GenericControllerAdapter[T]) Handle(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	resp, err := NewHttpResponse[T](w, r)
	if err != nil {
		return err
	}
	c.Controller.Handler(resp)

	return nil
}
