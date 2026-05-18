package core_http_server

import (
	"fmt"
	"net/http"
)

type Route struct {
	HandleFunc http.HandlerFunc
	Method string
	Url string
}

type Router struct {
	ServerMux *http.ServeMux
}

func NewRouter(mux *http.ServeMux) Router {
	return Router{
		ServerMux: mux,
	}
}

func (r *Router) RegisterRoutes(rts ...Route) {
	for _, v := range rts{
		r.ServerMux.Handle(fmt.Sprintf("%s %s", v.Method, v.Url), v.HandleFunc)
	}
}
