package api

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/ablqk/santorini/lib/errors"
)

// Endpoint defines a JSON endpoint.
type Endpoint interface {
	Path() string
	Verb() string
	NominalResponse() int
	Serve(r *http.Request) (Response, *errors.HTTPError)
}

// NewHandler creates a Handler for a given Endpoint.
func NewHandler(endpoint Endpoint) http.Handler {
	return handler{endpoint}
}

type handler struct {
	endpoint Endpoint
}

// ServeHTTP implementation of http.Handler.
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.endpoint.NominalResponse())

	resp, err := h.endpoint.Serve(r)
	if err != nil {
		fmt.Println("error while serving:", err)
		http.Error(w, err.Error(), err.GetErrorValue())
		return
	}

	js, errr := json.Marshal(resp)
	if errr != nil {
		fmt.Println("cannot marshal response:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, errr = w.Write(js)
	if errr != nil {
		fmt.Println("cannot write json:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
