package json

import (
	"encoding/json"
	"net/http"
)

type Presenter struct{}

func NewPresenter() *Presenter { return new(Presenter) }

func (p *Presenter) Present(rw http.ResponseWriter, statusCode int, data interface{}) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	return json.NewEncoder(rw).Encode(data)
}
