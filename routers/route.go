package routers

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
	Name    string
}
