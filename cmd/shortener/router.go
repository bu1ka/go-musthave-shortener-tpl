package main

import (
	"errors"
	"net/http"
)

type Router struct {
	routes []Route
}

type Pattern string

type Route struct {
	pattern Pattern
	method string
	action http.HandlerFunc
}

func NewRouter(routes []Route) Router {
	return Router{routes}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route, err := r.find(req)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Router not found"))
	} else {
		route.action(w, req)
	}
}

func (r *Router) find(req *http.Request) (Route, error) {
	for _, route := range r.routes {
		if route.method == req.Method {
			return route, nil
		}
	}

	return Route{}, errors.New("not found")
}