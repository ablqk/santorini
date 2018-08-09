package jsonhttp

import (
	"github.com/ablqk/santorini/lib/errors"
	"encoding/json"
	"net/http"
)

func DecodeBody(r *http.Request, o interface{}) *errors.HTTPError {
	err := json.NewDecoder(r.Body).Decode(o)
	if err != nil {
		return errors.New(errors.BadRequestE, "cannot decode body")
	}
	return nil
}
