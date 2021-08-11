package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error *JsonError  `json:"error"`
}

type JsonError struct {
	ErrorMessage string `json:"message"`
}

func (r *Response) Send(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(r); err != nil {
		log.Println(err)
	}
}

func NewResponse(data interface{}, err error) *Response {
	resp := &Response{
		Data: data,
	}

	if err != nil {
		resp.Error = &JsonError{
			fmt.Sprintf("%s", err),
		}
	}

	return resp
}

func sendHeaders(w http.ResponseWriter, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
}

func writeResponse(w http.ResponseWriter, msg string) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&JsonError{msg})

	if err != nil {
		log.Println(err)
	}
}

func NotFound(w http.ResponseWriter, msg string) {
	sendHeaders(w, http.StatusNotFound)
	writeResponse(w, msg)
}

func Unauthorized(w http.ResponseWriter, msg string) {
	sendHeaders(w, http.StatusUnauthorized)
	writeResponse(w, msg)
}

func Forbidden(w http.ResponseWriter, msg string) {
	sendHeaders(w, http.StatusForbidden)
	writeResponse(w, msg)
}

func InternalError(w http.ResponseWriter, msg string) {
	sendHeaders(w, http.StatusInternalServerError)
	writeResponse(w, msg)
}
