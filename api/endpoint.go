package api

import (
	"net/http"
	"encoding/json"
)

// Endpoint defines a JSON endpoint.
type Endpoint interface {
	Path() string
	Verb() string
	NominalResponse() int
	Serve(r *http.Request) (Response, error)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
