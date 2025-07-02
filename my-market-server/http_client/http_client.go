package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	port        string
	mux         *http.ServeMux
	controllers map[string]map[string]IController
}

type HttpException struct {
	Status     string `json:"status"`
	StatusCode int16  `json:"statusCode"`
	Message    string `json:"message"`
}

func (c *Client) RegisterController(basePath string, controllers ...IController) {

	for _, handle := range controllers {
		route := fmt.Sprintf("%s%s", basePath, handle.GetPath())
		method := strings.ToUpper(handle.GetMethod())
		if _, ok := c.controllers[route]; !ok {
			c.controllers[route] = map[string]IController{
				method: handle,
			}
			continue
		}

		if _, ok := c.controllers[route][method]; ok {
			panic(fmt.Sprintf("\033[31m[error] \033[0m[%s] %s is registered", strings.ToUpper(method), route))
		}

		c.controllers[route][method] = handle
	}
}

func NewClient() Client {
	mux := http.NewServeMux()

	return Client{
		mux:         mux,
		controllers: make(map[string]map[string]IController),
	}
}

func (c *Client) Listen(port string) {
	c.port = port

	c.serve()

	fmt.Printf("\033[32m[success] \033[0mApp successfully started on port %s\n", c.port)

	middleware := middleware(c.mux)

	err := http.ListenAndServe(fmt.Sprintf(":%s", c.port), middleware)

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
			if _, ok := controller[r.Method]; !ok {
				exp := HttpException{
					Status:     "error",
					StatusCode: http.StatusMethodNotAllowed,
					Message:    fmt.Sprintf("Method: [%s] %s not allowed", r.Method, r.Pattern),
				}

				data, _ := json.Marshal(exp)
				w.Header().Set(ct, ContentTypeJson)
				w.WriteHeader(int(exp.StatusCode))
				w.Write(data)
				return
			}

			ctrl := controller[r.Method]

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
				return
			}

		})
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
