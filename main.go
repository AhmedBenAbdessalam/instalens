package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content_Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>contact me here</h1>")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p><p>Q: Is there a free version?</p><p>A: Yes</p>")
}

func greeting(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	mood := r.URL.Query().Get("mood")
	fmt.Fprintf(w, "<h1>Greetings</h1> <p>hello %v you are feeling %v", name, mood)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Get("/", home)
	r.Get("/contact", contact)
	r.Get("/faq", faq)
	r.Get("/greeting/{name}", greeting)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port 3000....")
	http.ListenAndServe(":3000", r)
}
