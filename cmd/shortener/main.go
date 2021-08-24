package main

import (
	"io"
	"net/http"
)

var addr = "localhost:8080"

var dict Dict

func RootGetHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, err := getID(r.URL.Path)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)

		return
	}

	el, err := dict.get(id)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)

		return
	}

	http.Redirect(w, r, string(el), http.StatusTemporaryRedirect)
}

func RootPostHandler(w http.ResponseWriter, r *http.Request) {
	url, err := io.ReadAll(r.Body)

	if err != nil || len(url) == 0 {
		http.Error(w, "bad request", http.StatusBadRequest)

		return
	}

	short := dict.set(url)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(short))

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}



func main() {
	routes := []Route{
		{
			pattern: "/",
			method:  http.MethodGet,
			action:  RootGetHandler,
		},
		{
			pattern: "/",
			method:  http.MethodPost,
			action:  RootPostHandler,
		},
	}
	router := new(Router)
	router.routes = routes

	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic("server not started")
	}

	// Graceful shutdown
}
