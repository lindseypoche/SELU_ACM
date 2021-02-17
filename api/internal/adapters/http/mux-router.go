package router

// NOTE: Implementation complete but unused

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // access to mux library
// var (
// 	muxDispatcher = mux.NewRouter()
// )

// type muxRouter struct{}

// // NewMuxRouter creates a new mux router
// func NewMuxRouter() Router {
// 	return &muxRouter{}
// }

// func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	muxDispatcher.HandleFunc(uri, f).Methods("GET")
// }

// func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
// 	muxDispatcher.HandleFunc(uri, f).Methods("POST")
// }

// func (*muxRouter) SERVE(port string) {
// 	fmt.Printf("Mux HTTP server running on port %v", port)
// 	http.ListenAndServe(port, muxDispatcher)
// }
