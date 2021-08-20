package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

/**
Задание для трека «Go в веб-разработке»
Напишите сервис для сокращения длинных URL. Требования:
[x] Сервер должен быть доступен по адресу: http://localhost:8080.
[x] Сервер должен предоставлять два эндпоинта: POST / и GET /{id}.
[ ] Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает в ответ правильный сокращённый URL.
[ ] Эндпоинт GET /{id} принимает в качестве URL параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400.
 */

var addr = "http://localhost:8080"

type Dict struct {
	elems [][]byte
}

func (d *Dict) set(full []byte) string {
	d.elems = append(d.elems, full)

	return addr + "/" + strconv.Itoa(len(d.elems) - 1)
}

func (d *Dict) get(id int) ([]byte, error) {
	if len(d.elems) == 0 {
		return nil, errors.New("not found")
	}

	el := d.elems[id]

	fmt.Println("founded el", el)

	if len(el) != 0 {
		return el, nil
	}

	return nil, errors.New("not found")
}

var dict Dict

func RootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method == http.MethodPost {
		url, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)

			return
		}

		short := dict.set(url)

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte(short))

		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
		}

		return
	}

	if r.Method == http.MethodGet {
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


}

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(RootHandler))

	err := http.ListenAndServe(addr, router)

	if err != nil {
		panic("server not started")
	}

	// Graceful shutdown
}
