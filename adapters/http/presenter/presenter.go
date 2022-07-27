package presenter

import "net/http"

type Presenter interface {
	Present(rw http.ResponseWriter, statusCode int, data interface{}) error
}
