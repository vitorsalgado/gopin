package arch

import "net/http"

type Handler interface {
	Execute(w http.ResponseWriter, r *http.Request)
}
