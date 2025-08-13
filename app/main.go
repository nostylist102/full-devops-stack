//простой го файл
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from full-devops-stack!\n"))
	})
	http.ListenAndServe(":8080", r)
}