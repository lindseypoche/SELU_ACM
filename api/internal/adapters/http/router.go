package router

import "net/http"

// Router is an interface for http routers
// E.g. mux, chi
type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
