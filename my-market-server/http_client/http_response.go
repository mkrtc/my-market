package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
)

const ct string = "Content-type"
const ContentTypeJson string = "application/json"

type HttpResponse[T any] struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Query          url.Values
	ctx            context.Context
	Body           T
	Error          error
	HasBody        bool
}

type Exception struct {
	Status     string      `json:"status"`
	StatusCode int16       `json:"statusCode"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload"`
}

func NewHttpResponse[T any](rw http.ResponseWriter, rq *http.Request) (HttpResponse[T], error) {
	resp := HttpResponse[T]{
		ResponseWriter: rw,
		Request:        rq,
		Query:          rq.URL.Query(),
		ctx:            rq.Context(),
	}
	err := resp.setBody()
	resp.Error = err
	return resp, err
}

func (r *HttpResponse[T]) setBody() error {
	var body T
	err := json.NewDecoder(r.Request.Body).Decode(&body)
	if err != nil {
		r.HasBody = false
		if err.Error() == "EOF" {
			return nil
		}
		return err
	}
	r.HasBody = true

	validate := validator.New()

	err = validate.Struct(body)
	if err != nil {
		mess := ""
		for _, err := range err.(validator.ValidationErrors) {
			mess += fmt.Sprintf("%s;", err.Error())
		}
		return fmt.Errorf("%s", mess)
	}
	r.Body = body
	return nil
}

func (r *HttpResponse[T]) SendJson(data any) error {
	r.ResponseWriter.Header().Set(ct, ContentTypeJson)
	jsonData, parseJsonErr := json.Marshal(data)
	if parseJsonErr != nil {
		return parseJsonErr
	}

	_, err := r.ResponseWriter.Write(jsonData)

	return err

}

func (r *HttpResponse[T]) Exception(code int16, message string, payload any) error {
	exp := Exception{
		Status:     "error",
		StatusCode: code,
		Message:    message,
		Payload:    payload,
	}

	jsonExp, jsonExpErr := json.Marshal(exp)
	if jsonExpErr != nil {
		return jsonExpErr
	}

	r.ResponseWriter.Header().Set(ct, ContentTypeJson)
	r.ResponseWriter.WriteHeader(int(exp.StatusCode))
	_, wrErr := r.ResponseWriter.Write(jsonExp)

	return wrErr

}
