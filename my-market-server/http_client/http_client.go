package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
)

type Client struct {
	port        string
	mux         *http.ServeMux
	controllers map[string][]IController
}

type HttpException struct {
	Status     string `json:"status"`
	StatusCode int16  `json:"statusCode"`
	Message    string `json:"message"`
}

func (c *Client) RegisterController(basePath string, controllers ...IController) {

	for _, handle := range controllers {
		route := fmt.Sprintf("%s%s", basePath, handle.GetPath())

		if _, ok := c.controllers[route]; ok {
			rt := c.controllers[route]

			if slices.ContainsFunc(rt, func(con IController) bool {
				return con.GetMethod() == handle.GetMethod()
			}) {
				panic(fmt.Sprintf("\033[31m[error] \033[0m duplicate route: [%s] %s%s", handle.GetMethod(), basePath, handle.GetPath()))
			}
		}

		c.controllers[route] = append(c.controllers[route], handle)
	}
}

func NewClient() *Client {
	mux := http.NewServeMux()

	return &Client{
		mux:         mux,
		controllers: make(map[string][]IController),
	}
}

func (c *Client) Listen(port string) {
	c.port = port

	c.serve()

	fmt.Printf("\033[32m[success] \033[0mApp successfully started on port %s\n", c.port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", c.port), c.mux)

	if err != nil {
		panic(fmt.Sprintf("\033[31m[error] \033[0mCannot up server on port %s: %v\n", c.port, err))
	}
}

func (c *Client) serve() {
	if len(c.controllers) == 0 {
		fmt.Printf("\033[33m[warning] \033[0mno controller registered \n")
		return
	}
	for route, controller := range c.controllers {
		fmt.Printf("\033[32m[controller] \033[33m[%s] \033[0m successfully registered \n", route)
		c.mux.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			hasRoute := false
			var ctrl IController

			for _, conn := range controller {
				if r.Method == conn.GetMethod() {
					ctrl = conn
					hasRoute = true
				}
			}

			if hasRoute {
				ctx := r.Context()
				err := ctrl.Handle(ctx, w, r)
				if err != nil {
					exp := HttpException{
						Status:     "error",
						StatusCode: http.StatusBadRequest,
						Message:    err.Error(),
					}
					data, _ := json.Marshal(exp)
					w.Header().Set(ct, ContentTypeJson)
					w.WriteHeader(int(exp.StatusCode))
					w.Write(data)
				}
				return
			}

			exp := HttpException{
				Status:     "error",
				StatusCode: http.StatusMethodNotAllowed,
				Message:    fmt.Sprintf("Method: [%s] %s not allowed", r.Method, r.Pattern),
			}

			data, _ := json.Marshal(exp)
			w.Header().Set(ct, ContentTypeJson)
			w.WriteHeader(int(exp.StatusCode))
			w.Write(data)

		})
	}
}
