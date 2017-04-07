package httpHelper

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	HTTPRequest *http.Request
	Params      url.Values
}

func NewRequest(req *http.Request, params url.Values) Request {
	return Request{req, params}
}

func (r *Request) ParseJSONBody(obj interface{}) (interface{}, error) {
	if err := json.NewDecoder(r.HTTPRequest.Body).Decode(obj); err != nil {
		return nil, err
	}
	return obj, nil
}
