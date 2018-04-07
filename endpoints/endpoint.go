package endpoints

import (
	"net/http"
	"encoding/json"
	"github.com/ablqk/santorini/definitions"
)

// Endpoint defines a JSON endpoint.
type Endpoint interface {
	Path() string
	Verb() string
	NominalResponse() int
	Serve(r *http.Request) (definitions.Response, error)
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
	}

	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
