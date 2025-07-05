package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func plusNums(w http.ResponseWriter, r *http.Request) {
	n1 := chi.URLParam(r, "n1")
	n2 := chi.URLParam(r, "n2")
	num1, err := strconv.Atoi(n1)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(n2)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(num1 + num2)))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("chi router"))
	})

	r.Get("/plus/{n1}-{n2}", plusNums)

	http.ListenAndServe(":8080", r)
}
